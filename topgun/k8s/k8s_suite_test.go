package k8s_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"
	"time"

	"github.com/caarlos0/env"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	. "github.com/concourse/concourse/topgun"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestK8s(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "K8s Suite")
}

type environment struct {
	HelmChartsDir        string `env:"HELM_CHARTS_DIR,required"`
	ConcourseChartDir    string `env:"CONCOURSE_CHART_DIR,required"`
	ConcourseImageDigest string `env:"CONCOURSE_IMAGE_DIGEST"`
	ConcourseImageName   string `env:"CONCOURSE_IMAGE_NAME,required"`
	ConcourseImageTag    string `env:"CONCOURSE_IMAGE_TAG"`
	FlyPath              string `env:"FLY_PATH"`
	K8sEngine            string `env:"K8S_ENGINE" envDefault:"GKE"`
	InCluster            bool   `env:"IN_CLUSTER" envDefault:"false"`
}

var (
	Environment            environment
	endpointFactory        EndpointFactory
	fly                    FlyCli
	releaseName, namespace string
)

var _ = SynchronizedBeforeSuite(func() []byte {
	var parsedEnv environment

	err := env.Parse(&parsedEnv)
	Expect(err).ToNot(HaveOccurred())

	if parsedEnv.FlyPath == "" {
		parsedEnv.FlyPath = BuildBinary()
	}

	By("Checking if kubectl has a context set")
	Wait(Start(nil, "kubectl", "config", "current-context"))

	By("Initializing the client side of helm")
	Wait(Start(nil, "helm", "init", "--client-only"))

	By("Updating the dependencies of the Concourse chart locally")
	Wait(Start(nil, "helm", "dependency", "update", parsedEnv.ConcourseChartDir))

	envBytes, err := json.Marshal(parsedEnv)
	Expect(err).ToNot(HaveOccurred())

	return envBytes
}, func(data []byte) {
	err := json.Unmarshal(data, &Environment)
	Expect(err).ToNot(HaveOccurred())
})

var _ = BeforeEach(func() {
	SetDefaultEventuallyTimeout(90 * time.Second)
	SetDefaultConsistentlyDuration(30 * time.Second)

	tmp, err := ioutil.TempDir("", "topgun-tmp")
	Expect(err).ToNot(HaveOccurred())

	fly = FlyCli{
		Bin:    Environment.FlyPath,
		Target: "concourse-topgun-k8s-" + strconv.Itoa(GinkgoParallelNode()),
		Home:   filepath.Join(tmp, "fly-home-"+strconv.Itoa(GinkgoParallelNode())),
	}

	endpointFactory = PortForwardingEndpointFactory{}
	if Environment.InCluster {
		endpointFactory = AddressEndpointFactory{}
	}

	err = os.Mkdir(fly.Home, 0755)
	Expect(err).ToNot(HaveOccurred())
})

func setReleaseNameAndNamespace(description string) {
	rand.Seed(time.Now().UTC().UnixNano())
	releaseName = fmt.Sprintf("topgun-"+description+"-%d", rand.Int63n(100000000))
	namespace = releaseName
}

// pod corresponds to the Json object that represents a Kuberneted pod from the
// apiserver perspective.
//
type pod struct {
	Status struct {
		Conditions []struct {
			Type   string `json:"type"`
			Status string `json:"status"`
		} `json:"conditions"`
		ContainerStatuses []struct {
			Name  string `json:"name"`
			Ready bool   `json:"ready"`
		} `json:"containerStatuses"`
		Ip string `json:"podIP"`
	} `json:"status"`
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
}

// Endpoint represents a service that can be reached from a given address.
//
type Endpoint interface {
	Address() (addr string)
	Close() (err error)
}

// EndpointFactory represents those entities able to generate Endpoints for
// both services and pods.
//
type EndpointFactory interface {
	NewServiceEndpoint(namespace, service, port string) (endpoint Endpoint)
	NewPodEndpoint(namespace, pod, port string) (endpoint Endpoint)
}

// PortForwardingEndpoint is a service that can be reached through a local
// address, having connections port forwarded to entities in a cluster.
//
type PortForwardingEndpoint struct {
	session *gexec.Session
	address string
}

func (p PortForwardingEndpoint) Address() string {
	return p.address
}

func (p PortForwardingEndpoint) Close() error {
	p.session.Interrupt()
	return nil
}

// AddressEndpoint represents a direct address without any underlying session.
//
type AddressEndpoint struct {
	address string
}

func (p AddressEndpoint) Address() string {
	return p.address
}

func (p AddressEndpoint) Close() error {
	return nil
}

// PortForwardingFactory deals with creating endpoints that reach the targets
// through port-forwarding.
//
type PortForwardingEndpointFactory struct{}

func (f PortForwardingEndpointFactory) NewServiceEndpoint(namespace, service, port string) Endpoint {
	session, address := portForward(namespace, "service/"+service, port)

	return PortForwardingEndpoint{
		session: session,
		address: address,
	}
}

func (f PortForwardingEndpointFactory) NewPodEndpoint(namespace, pod, port string) Endpoint {
	session, address := portForward(namespace, "pod/"+pod, port)

	return PortForwardingEndpoint{
		session: session,
		address: address,
	}
}

// AddressFactory deals with creating endpoints that reach the targets
// through port-forwarding.
//
type AddressEndpointFactory struct{}

func (f AddressEndpointFactory) NewServiceEndpoint(namespace, service, port string) Endpoint {
	address := serviceAddress(namespace, service)

	return AddressEndpoint{
		address: address + ":" + port,
	}
}

func (f AddressEndpointFactory) NewPodEndpoint(namespace, pod, port string) Endpoint {
	address := podAddress(namespace, pod)

	return AddressEndpoint{
		address: address + ":" + port,
	}
}

func podAddress(namespace, pod string) (address string) {
	pods := getPods(namespace, "--field-selector=metadata.name="+pod)
	Expect(pods).To(HaveLen(1))

	return pods[0].Status.Ip
}

// serviceAddress retrieves the ClusterIP address of a service on a given
// namespace.
//
func serviceAddress(namespace, serviceName string) (address string) {
	return serviceName + "." + namespace
}

// portForward establishes a port-forwarding session against a given kubernetes
// resource, for a particular port.
//
func portForward(namespace, resource, port string) (*gexec.Session, string) {
	sess := Start(nil,
		"kubectl", "port-forward",
		"--namespace="+namespace,
		resource,
		":"+port,
	)

	Eventually(sess.Out).Should(gbytes.Say("Forwarding"))

	address := regexp.MustCompile(`127\.0\.0\.1:[0-9]+`).
		FindStringSubmatch(string(sess.Out.Contents()))
	Expect(address).NotTo(BeEmpty())

	return sess, address[0]
}

func helmDeploy(releaseName, namespace, chartDir string, args ...string) *gexec.Session {
	helmArgs := []string{
		"upgrade",
		"--install",
		"--force",
		"--namespace", namespace,
	}

	helmArgs = append(helmArgs, args...)
	helmArgs = append(helmArgs, releaseName, chartDir)

	sess := Start(nil, "helm", helmArgs...)
	<-sess.Exited
	return sess
}

func helmInstallArgs(args ...string) []string {
	helmArgs := []string{
		"--set=concourse.web.kubernetes.keepNamespaces=false",
		"--set=concourse.worker.bindIp=0.0.0.0",
		"--set=postgresql.persistence.enabled=false",
		"--set=web.resources.requests.cpu=500m",
		"--set=worker.readinessProbe.httpGet.path=/",
		"--set=worker.readinessProbe.httpGet.port=worker-hc",
		"--set=worker.resources.requests.cpu=500m",
		"--set=image=" + Environment.ConcourseImageName}

	if Environment.ConcourseImageTag != "" {
		helmArgs = append(helmArgs, "--set=imageTag="+Environment.ConcourseImageTag)
	}

	if Environment.ConcourseImageDigest != "" {
		helmArgs = append(helmArgs, "--set=imageDigest="+Environment.ConcourseImageDigest)
	}

	return append(helmArgs, args...)
}

func deployFailingConcourseChart(releaseName string, expectedErr string, args ...string) {
	helmArgs := helmInstallArgs(args...)
	sess := helmDeploy(releaseName, releaseName, Environment.ConcourseChartDir, helmArgs...)
	Expect(sess.ExitCode()).ToNot(Equal(0))
	Expect(sess.Err).To(gbytes.Say(expectedErr))
}

func deployConcourseChart(releaseName string, args ...string) {
	helmArgs := helmInstallArgs(args...)
	sess := helmDeploy(releaseName, releaseName, Environment.ConcourseChartDir, helmArgs...)
	Expect(sess.ExitCode()).To(Equal(0))
}

func helmDestroy(releaseName string) {
	helmArgs := []string{
		"delete",
		"--purge",
		releaseName,
	}

	Wait(Start(nil, "helm", helmArgs...))
}

func getPods(namespace string, flags ...string) []pod {
	var (
		pods struct {
			Items []pod `json:"items"`
		}

		args = append([]string{"get", "pods",
			"--namespace=" + namespace,
			"--output=json",
			"--no-headers"}, flags...)
		session = Start(nil, "kubectl", args...)
	)

	Wait(session)

	err := json.Unmarshal(session.Out.Contents(), &pods)
	Expect(err).ToNot(HaveOccurred())

	return pods.Items
}

func isPodReady(p pod) bool {
	for _, condition := range p.Status.Conditions {
		if condition.Type != "ContainersReady" {
			continue
		}

		return condition.Status == "True"
	}

	return false
}

func waitAllPodsInNamespaceToBeReady(namespace string) {
	Eventually(func() bool {
		expectedPods := getPods(namespace)
		actualPods := getPods(namespace, "--field-selector=status.phase=Running")

		if len(expectedPods) != len(actualPods) {
			return false
		}

		podsReady := 0
		for _, pod := range actualPods {
			if isPodReady(pod) {
				podsReady++
			}
		}

		return podsReady == len(expectedPods)
	}, 15*time.Minute, 10*time.Second).Should(BeTrue(), "expected all pods to be running")
}

func deletePods(namespace string, flags ...string) []string {
	var (
		podNames []string
		args     = append([]string{"delete", "pod",
			"--namespace=" + namespace,
		}, flags...)
		session = Start(nil, "kubectl", args...)
	)

	Wait(session)

	scanner := bufio.NewScanner(bytes.NewBuffer(session.Out.Contents()))
	for scanner.Scan() {
		podNames = append(podNames, scanner.Text())
	}

	return podNames
}

func getRunningWorkers(workers []Worker) (running []Worker) {
	for _, w := range workers {
		if w.State == "running" {
			running = append(running, w)
		}
	}
	return
}

func waitAndLogin(namespace, service string) Endpoint {
	waitAllPodsInNamespaceToBeReady(namespace)

	atc := endpointFactory.NewServiceEndpoint(
		namespace,
		service,
		"8080",
	)

	fly.Login("test", "test", "http://"+atc.Address())

	Eventually(func() []Worker {
		return getRunningWorkers(fly.GetWorkers())
	}, 2*time.Minute, 10*time.Second).
		ShouldNot(HaveLen(0))

	return atc
}

func cleanup(releaseName, namespace string) {
	helmDestroy(releaseName)
	Run(nil, "kubectl", "delete", "namespace", namespace, "--wait=false")
}

func onPks(f func()) {
	Context("PKS", func() {

		BeforeEach(func() {
			if Environment.K8sEngine != "PKS" {
				Skip("not running on PKS")
			}
		})

		f()
	})
}

func onGke(f func()) {
	Context("GKE", func() {

		BeforeEach(func() {
			if Environment.K8sEngine != "GKE" {
				Skip("not running on GKE")
			}
		})

		f()
	})
}
