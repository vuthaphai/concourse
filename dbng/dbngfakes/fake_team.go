// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

type FakeTeam struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	SavePipelineStub        func(pipelineName string, config atc.Config, from dbng.ConfigVersion, pausedState dbng.PipelinePausedState) (dbng.Pipeline, bool, error)
	savePipelineMutex       sync.RWMutex
	savePipelineArgsForCall []struct {
		pipelineName string
		config       atc.Config
		from         dbng.ConfigVersion
		pausedState  dbng.PipelinePausedState
	}
	savePipelineReturns struct {
		result1 dbng.Pipeline
		result2 bool
		result3 error
	}
	FindPipelineByNameStub        func(pipelineName string) (dbng.Pipeline, bool, error)
	findPipelineByNameMutex       sync.RWMutex
	findPipelineByNameArgsForCall []struct {
		pipelineName string
	}
	findPipelineByNameReturns struct {
		result1 dbng.Pipeline
		result2 bool
		result3 error
	}
	CreateOneOffBuildStub        func() (dbng.Build, error)
	createOneOffBuildMutex       sync.RWMutex
	createOneOffBuildArgsForCall []struct{}
	createOneOffBuildReturns     struct {
		result1 dbng.Build
		result2 error
	}
	SaveWorkerStub        func(atcWorker atc.Worker, ttl time.Duration) (dbng.Worker, error)
	saveWorkerMutex       sync.RWMutex
	saveWorkerArgsForCall []struct {
		atcWorker atc.Worker
		ttl       time.Duration
	}
	saveWorkerReturns struct {
		result1 dbng.Worker
		result2 error
	}
	WorkersStub        func() ([]dbng.Worker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct{}
	workersReturns     struct {
		result1 []dbng.Worker
		result2 error
	}
	FindContainerByHandleStub        func(string) (dbng.CreatedContainer, bool, error)
	findContainerByHandleMutex       sync.RWMutex
	findContainerByHandleArgsForCall []struct {
		arg1 string
	}
	findContainerByHandleReturns struct {
		result1 dbng.CreatedContainer
		result2 bool
		result3 error
	}
	FindWorkerForResourceCheckContainerStub        func(resourceConfig *dbng.UsedResourceConfig) (dbng.Worker, bool, error)
	findWorkerForResourceCheckContainerMutex       sync.RWMutex
	findWorkerForResourceCheckContainerArgsForCall []struct {
		resourceConfig *dbng.UsedResourceConfig
	}
	findWorkerForResourceCheckContainerReturns struct {
		result1 dbng.Worker
		result2 bool
		result3 error
	}
	FindResourceCheckContainerOnWorkerStub        func(workerName string, resourceConfig *dbng.UsedResourceConfig) (dbng.CreatingContainer, dbng.CreatedContainer, error)
	findResourceCheckContainerOnWorkerMutex       sync.RWMutex
	findResourceCheckContainerOnWorkerArgsForCall []struct {
		workerName     string
		resourceConfig *dbng.UsedResourceConfig
	}
	findResourceCheckContainerOnWorkerReturns struct {
		result1 dbng.CreatingContainer
		result2 dbng.CreatedContainer
		result3 error
	}
	CreateResourceCheckContainerStub        func(workerName string, resourceConfig *dbng.UsedResourceConfig) (dbng.CreatingContainer, error)
	createResourceCheckContainerMutex       sync.RWMutex
	createResourceCheckContainerArgsForCall []struct {
		workerName     string
		resourceConfig *dbng.UsedResourceConfig
	}
	createResourceCheckContainerReturns struct {
		result1 dbng.CreatingContainer
		result2 error
	}
	CreateResourceGetContainerStub        func(workerName string, resourceConfig *dbng.UsedResourceCache, stepName string) (dbng.CreatingContainer, error)
	createResourceGetContainerMutex       sync.RWMutex
	createResourceGetContainerArgsForCall []struct {
		workerName     string
		resourceConfig *dbng.UsedResourceCache
		stepName       string
	}
	createResourceGetContainerReturns struct {
		result1 dbng.CreatingContainer
		result2 error
	}
	FindWorkerForBuildContainerStub        func(buildID int, planID atc.PlanID) (dbng.Worker, bool, error)
	findWorkerForBuildContainerMutex       sync.RWMutex
	findWorkerForBuildContainerArgsForCall []struct {
		buildID int
		planID  atc.PlanID
	}
	findWorkerForBuildContainerReturns struct {
		result1 dbng.Worker
		result2 bool
		result3 error
	}
	FindBuildContainerOnWorkerStub        func(workerName string, buildID int, planID atc.PlanID) (dbng.CreatingContainer, dbng.CreatedContainer, error)
	findBuildContainerOnWorkerMutex       sync.RWMutex
	findBuildContainerOnWorkerArgsForCall []struct {
		workerName string
		buildID    int
		planID     atc.PlanID
	}
	findBuildContainerOnWorkerReturns struct {
		result1 dbng.CreatingContainer
		result2 dbng.CreatedContainer
		result3 error
	}
	CreateBuildContainerStub        func(workerName string, buildID int, planID atc.PlanID, meta dbng.ContainerMetadata) (dbng.CreatingContainer, error)
	createBuildContainerMutex       sync.RWMutex
	createBuildContainerArgsForCall []struct {
		workerName string
		buildID    int
		planID     atc.PlanID
		meta       dbng.ContainerMetadata
	}
	createBuildContainerReturns struct {
		result1 dbng.CreatingContainer
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeam) ID() int {
	fake.iDMutex.Lock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	return fake.iDReturns.result1
}

func (fake *FakeTeam) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeTeam) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeTeam) SavePipeline(pipelineName string, config atc.Config, from dbng.ConfigVersion, pausedState dbng.PipelinePausedState) (dbng.Pipeline, bool, error) {
	fake.savePipelineMutex.Lock()
	fake.savePipelineArgsForCall = append(fake.savePipelineArgsForCall, struct {
		pipelineName string
		config       atc.Config
		from         dbng.ConfigVersion
		pausedState  dbng.PipelinePausedState
	}{pipelineName, config, from, pausedState})
	fake.recordInvocation("SavePipeline", []interface{}{pipelineName, config, from, pausedState})
	fake.savePipelineMutex.Unlock()
	if fake.SavePipelineStub != nil {
		return fake.SavePipelineStub(pipelineName, config, from, pausedState)
	}
	return fake.savePipelineReturns.result1, fake.savePipelineReturns.result2, fake.savePipelineReturns.result3
}

func (fake *FakeTeam) SavePipelineCallCount() int {
	fake.savePipelineMutex.RLock()
	defer fake.savePipelineMutex.RUnlock()
	return len(fake.savePipelineArgsForCall)
}

func (fake *FakeTeam) SavePipelineArgsForCall(i int) (string, atc.Config, dbng.ConfigVersion, dbng.PipelinePausedState) {
	fake.savePipelineMutex.RLock()
	defer fake.savePipelineMutex.RUnlock()
	return fake.savePipelineArgsForCall[i].pipelineName, fake.savePipelineArgsForCall[i].config, fake.savePipelineArgsForCall[i].from, fake.savePipelineArgsForCall[i].pausedState
}

func (fake *FakeTeam) SavePipelineReturns(result1 dbng.Pipeline, result2 bool, result3 error) {
	fake.SavePipelineStub = nil
	fake.savePipelineReturns = struct {
		result1 dbng.Pipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) FindPipelineByName(pipelineName string) (dbng.Pipeline, bool, error) {
	fake.findPipelineByNameMutex.Lock()
	fake.findPipelineByNameArgsForCall = append(fake.findPipelineByNameArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.recordInvocation("FindPipelineByName", []interface{}{pipelineName})
	fake.findPipelineByNameMutex.Unlock()
	if fake.FindPipelineByNameStub != nil {
		return fake.FindPipelineByNameStub(pipelineName)
	}
	return fake.findPipelineByNameReturns.result1, fake.findPipelineByNameReturns.result2, fake.findPipelineByNameReturns.result3
}

func (fake *FakeTeam) FindPipelineByNameCallCount() int {
	fake.findPipelineByNameMutex.RLock()
	defer fake.findPipelineByNameMutex.RUnlock()
	return len(fake.findPipelineByNameArgsForCall)
}

func (fake *FakeTeam) FindPipelineByNameArgsForCall(i int) string {
	fake.findPipelineByNameMutex.RLock()
	defer fake.findPipelineByNameMutex.RUnlock()
	return fake.findPipelineByNameArgsForCall[i].pipelineName
}

func (fake *FakeTeam) FindPipelineByNameReturns(result1 dbng.Pipeline, result2 bool, result3 error) {
	fake.FindPipelineByNameStub = nil
	fake.findPipelineByNameReturns = struct {
		result1 dbng.Pipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) CreateOneOffBuild() (dbng.Build, error) {
	fake.createOneOffBuildMutex.Lock()
	fake.createOneOffBuildArgsForCall = append(fake.createOneOffBuildArgsForCall, struct{}{})
	fake.recordInvocation("CreateOneOffBuild", []interface{}{})
	fake.createOneOffBuildMutex.Unlock()
	if fake.CreateOneOffBuildStub != nil {
		return fake.CreateOneOffBuildStub()
	}
	return fake.createOneOffBuildReturns.result1, fake.createOneOffBuildReturns.result2
}

func (fake *FakeTeam) CreateOneOffBuildCallCount() int {
	fake.createOneOffBuildMutex.RLock()
	defer fake.createOneOffBuildMutex.RUnlock()
	return len(fake.createOneOffBuildArgsForCall)
}

func (fake *FakeTeam) CreateOneOffBuildReturns(result1 dbng.Build, result2 error) {
	fake.CreateOneOffBuildStub = nil
	fake.createOneOffBuildReturns = struct {
		result1 dbng.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeTeam) SaveWorker(atcWorker atc.Worker, ttl time.Duration) (dbng.Worker, error) {
	fake.saveWorkerMutex.Lock()
	fake.saveWorkerArgsForCall = append(fake.saveWorkerArgsForCall, struct {
		atcWorker atc.Worker
		ttl       time.Duration
	}{atcWorker, ttl})
	fake.recordInvocation("SaveWorker", []interface{}{atcWorker, ttl})
	fake.saveWorkerMutex.Unlock()
	if fake.SaveWorkerStub != nil {
		return fake.SaveWorkerStub(atcWorker, ttl)
	}
	return fake.saveWorkerReturns.result1, fake.saveWorkerReturns.result2
}

func (fake *FakeTeam) SaveWorkerCallCount() int {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return len(fake.saveWorkerArgsForCall)
}

func (fake *FakeTeam) SaveWorkerArgsForCall(i int) (atc.Worker, time.Duration) {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return fake.saveWorkerArgsForCall[i].atcWorker, fake.saveWorkerArgsForCall[i].ttl
}

func (fake *FakeTeam) SaveWorkerReturns(result1 dbng.Worker, result2 error) {
	fake.SaveWorkerStub = nil
	fake.saveWorkerReturns = struct {
		result1 dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeTeam) Workers() ([]dbng.Worker, error) {
	fake.workersMutex.Lock()
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct{}{})
	fake.recordInvocation("Workers", []interface{}{})
	fake.workersMutex.Unlock()
	if fake.WorkersStub != nil {
		return fake.WorkersStub()
	}
	return fake.workersReturns.result1, fake.workersReturns.result2
}

func (fake *FakeTeam) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeTeam) WorkersReturns(result1 []dbng.Worker, result2 error) {
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeTeam) FindContainerByHandle(arg1 string) (dbng.CreatedContainer, bool, error) {
	fake.findContainerByHandleMutex.Lock()
	fake.findContainerByHandleArgsForCall = append(fake.findContainerByHandleArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FindContainerByHandle", []interface{}{arg1})
	fake.findContainerByHandleMutex.Unlock()
	if fake.FindContainerByHandleStub != nil {
		return fake.FindContainerByHandleStub(arg1)
	}
	return fake.findContainerByHandleReturns.result1, fake.findContainerByHandleReturns.result2, fake.findContainerByHandleReturns.result3
}

func (fake *FakeTeam) FindContainerByHandleCallCount() int {
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	return len(fake.findContainerByHandleArgsForCall)
}

func (fake *FakeTeam) FindContainerByHandleArgsForCall(i int) string {
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	return fake.findContainerByHandleArgsForCall[i].arg1
}

func (fake *FakeTeam) FindContainerByHandleReturns(result1 dbng.CreatedContainer, result2 bool, result3 error) {
	fake.FindContainerByHandleStub = nil
	fake.findContainerByHandleReturns = struct {
		result1 dbng.CreatedContainer
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) FindWorkerForResourceCheckContainer(resourceConfig *dbng.UsedResourceConfig) (dbng.Worker, bool, error) {
	fake.findWorkerForResourceCheckContainerMutex.Lock()
	fake.findWorkerForResourceCheckContainerArgsForCall = append(fake.findWorkerForResourceCheckContainerArgsForCall, struct {
		resourceConfig *dbng.UsedResourceConfig
	}{resourceConfig})
	fake.recordInvocation("FindWorkerForResourceCheckContainer", []interface{}{resourceConfig})
	fake.findWorkerForResourceCheckContainerMutex.Unlock()
	if fake.FindWorkerForResourceCheckContainerStub != nil {
		return fake.FindWorkerForResourceCheckContainerStub(resourceConfig)
	}
	return fake.findWorkerForResourceCheckContainerReturns.result1, fake.findWorkerForResourceCheckContainerReturns.result2, fake.findWorkerForResourceCheckContainerReturns.result3
}

func (fake *FakeTeam) FindWorkerForResourceCheckContainerCallCount() int {
	fake.findWorkerForResourceCheckContainerMutex.RLock()
	defer fake.findWorkerForResourceCheckContainerMutex.RUnlock()
	return len(fake.findWorkerForResourceCheckContainerArgsForCall)
}

func (fake *FakeTeam) FindWorkerForResourceCheckContainerArgsForCall(i int) *dbng.UsedResourceConfig {
	fake.findWorkerForResourceCheckContainerMutex.RLock()
	defer fake.findWorkerForResourceCheckContainerMutex.RUnlock()
	return fake.findWorkerForResourceCheckContainerArgsForCall[i].resourceConfig
}

func (fake *FakeTeam) FindWorkerForResourceCheckContainerReturns(result1 dbng.Worker, result2 bool, result3 error) {
	fake.FindWorkerForResourceCheckContainerStub = nil
	fake.findWorkerForResourceCheckContainerReturns = struct {
		result1 dbng.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) FindResourceCheckContainerOnWorker(workerName string, resourceConfig *dbng.UsedResourceConfig) (dbng.CreatingContainer, dbng.CreatedContainer, error) {
	fake.findResourceCheckContainerOnWorkerMutex.Lock()
	fake.findResourceCheckContainerOnWorkerArgsForCall = append(fake.findResourceCheckContainerOnWorkerArgsForCall, struct {
		workerName     string
		resourceConfig *dbng.UsedResourceConfig
	}{workerName, resourceConfig})
	fake.recordInvocation("FindResourceCheckContainerOnWorker", []interface{}{workerName, resourceConfig})
	fake.findResourceCheckContainerOnWorkerMutex.Unlock()
	if fake.FindResourceCheckContainerOnWorkerStub != nil {
		return fake.FindResourceCheckContainerOnWorkerStub(workerName, resourceConfig)
	}
	return fake.findResourceCheckContainerOnWorkerReturns.result1, fake.findResourceCheckContainerOnWorkerReturns.result2, fake.findResourceCheckContainerOnWorkerReturns.result3
}

func (fake *FakeTeam) FindResourceCheckContainerOnWorkerCallCount() int {
	fake.findResourceCheckContainerOnWorkerMutex.RLock()
	defer fake.findResourceCheckContainerOnWorkerMutex.RUnlock()
	return len(fake.findResourceCheckContainerOnWorkerArgsForCall)
}

func (fake *FakeTeam) FindResourceCheckContainerOnWorkerArgsForCall(i int) (string, *dbng.UsedResourceConfig) {
	fake.findResourceCheckContainerOnWorkerMutex.RLock()
	defer fake.findResourceCheckContainerOnWorkerMutex.RUnlock()
	return fake.findResourceCheckContainerOnWorkerArgsForCall[i].workerName, fake.findResourceCheckContainerOnWorkerArgsForCall[i].resourceConfig
}

func (fake *FakeTeam) FindResourceCheckContainerOnWorkerReturns(result1 dbng.CreatingContainer, result2 dbng.CreatedContainer, result3 error) {
	fake.FindResourceCheckContainerOnWorkerStub = nil
	fake.findResourceCheckContainerOnWorkerReturns = struct {
		result1 dbng.CreatingContainer
		result2 dbng.CreatedContainer
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) CreateResourceCheckContainer(workerName string, resourceConfig *dbng.UsedResourceConfig) (dbng.CreatingContainer, error) {
	fake.createResourceCheckContainerMutex.Lock()
	fake.createResourceCheckContainerArgsForCall = append(fake.createResourceCheckContainerArgsForCall, struct {
		workerName     string
		resourceConfig *dbng.UsedResourceConfig
	}{workerName, resourceConfig})
	fake.recordInvocation("CreateResourceCheckContainer", []interface{}{workerName, resourceConfig})
	fake.createResourceCheckContainerMutex.Unlock()
	if fake.CreateResourceCheckContainerStub != nil {
		return fake.CreateResourceCheckContainerStub(workerName, resourceConfig)
	}
	return fake.createResourceCheckContainerReturns.result1, fake.createResourceCheckContainerReturns.result2
}

func (fake *FakeTeam) CreateResourceCheckContainerCallCount() int {
	fake.createResourceCheckContainerMutex.RLock()
	defer fake.createResourceCheckContainerMutex.RUnlock()
	return len(fake.createResourceCheckContainerArgsForCall)
}

func (fake *FakeTeam) CreateResourceCheckContainerArgsForCall(i int) (string, *dbng.UsedResourceConfig) {
	fake.createResourceCheckContainerMutex.RLock()
	defer fake.createResourceCheckContainerMutex.RUnlock()
	return fake.createResourceCheckContainerArgsForCall[i].workerName, fake.createResourceCheckContainerArgsForCall[i].resourceConfig
}

func (fake *FakeTeam) CreateResourceCheckContainerReturns(result1 dbng.CreatingContainer, result2 error) {
	fake.CreateResourceCheckContainerStub = nil
	fake.createResourceCheckContainerReturns = struct {
		result1 dbng.CreatingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeTeam) CreateResourceGetContainer(workerName string, resourceConfig *dbng.UsedResourceCache, stepName string) (dbng.CreatingContainer, error) {
	fake.createResourceGetContainerMutex.Lock()
	fake.createResourceGetContainerArgsForCall = append(fake.createResourceGetContainerArgsForCall, struct {
		workerName     string
		resourceConfig *dbng.UsedResourceCache
		stepName       string
	}{workerName, resourceConfig, stepName})
	fake.recordInvocation("CreateResourceGetContainer", []interface{}{workerName, resourceConfig, stepName})
	fake.createResourceGetContainerMutex.Unlock()
	if fake.CreateResourceGetContainerStub != nil {
		return fake.CreateResourceGetContainerStub(workerName, resourceConfig, stepName)
	}
	return fake.createResourceGetContainerReturns.result1, fake.createResourceGetContainerReturns.result2
}

func (fake *FakeTeam) CreateResourceGetContainerCallCount() int {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return len(fake.createResourceGetContainerArgsForCall)
}

func (fake *FakeTeam) CreateResourceGetContainerArgsForCall(i int) (string, *dbng.UsedResourceCache, string) {
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	return fake.createResourceGetContainerArgsForCall[i].workerName, fake.createResourceGetContainerArgsForCall[i].resourceConfig, fake.createResourceGetContainerArgsForCall[i].stepName
}

func (fake *FakeTeam) CreateResourceGetContainerReturns(result1 dbng.CreatingContainer, result2 error) {
	fake.CreateResourceGetContainerStub = nil
	fake.createResourceGetContainerReturns = struct {
		result1 dbng.CreatingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeTeam) FindWorkerForBuildContainer(buildID int, planID atc.PlanID) (dbng.Worker, bool, error) {
	fake.findWorkerForBuildContainerMutex.Lock()
	fake.findWorkerForBuildContainerArgsForCall = append(fake.findWorkerForBuildContainerArgsForCall, struct {
		buildID int
		planID  atc.PlanID
	}{buildID, planID})
	fake.recordInvocation("FindWorkerForBuildContainer", []interface{}{buildID, planID})
	fake.findWorkerForBuildContainerMutex.Unlock()
	if fake.FindWorkerForBuildContainerStub != nil {
		return fake.FindWorkerForBuildContainerStub(buildID, planID)
	}
	return fake.findWorkerForBuildContainerReturns.result1, fake.findWorkerForBuildContainerReturns.result2, fake.findWorkerForBuildContainerReturns.result3
}

func (fake *FakeTeam) FindWorkerForBuildContainerCallCount() int {
	fake.findWorkerForBuildContainerMutex.RLock()
	defer fake.findWorkerForBuildContainerMutex.RUnlock()
	return len(fake.findWorkerForBuildContainerArgsForCall)
}

func (fake *FakeTeam) FindWorkerForBuildContainerArgsForCall(i int) (int, atc.PlanID) {
	fake.findWorkerForBuildContainerMutex.RLock()
	defer fake.findWorkerForBuildContainerMutex.RUnlock()
	return fake.findWorkerForBuildContainerArgsForCall[i].buildID, fake.findWorkerForBuildContainerArgsForCall[i].planID
}

func (fake *FakeTeam) FindWorkerForBuildContainerReturns(result1 dbng.Worker, result2 bool, result3 error) {
	fake.FindWorkerForBuildContainerStub = nil
	fake.findWorkerForBuildContainerReturns = struct {
		result1 dbng.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) FindBuildContainerOnWorker(workerName string, buildID int, planID atc.PlanID) (dbng.CreatingContainer, dbng.CreatedContainer, error) {
	fake.findBuildContainerOnWorkerMutex.Lock()
	fake.findBuildContainerOnWorkerArgsForCall = append(fake.findBuildContainerOnWorkerArgsForCall, struct {
		workerName string
		buildID    int
		planID     atc.PlanID
	}{workerName, buildID, planID})
	fake.recordInvocation("FindBuildContainerOnWorker", []interface{}{workerName, buildID, planID})
	fake.findBuildContainerOnWorkerMutex.Unlock()
	if fake.FindBuildContainerOnWorkerStub != nil {
		return fake.FindBuildContainerOnWorkerStub(workerName, buildID, planID)
	}
	return fake.findBuildContainerOnWorkerReturns.result1, fake.findBuildContainerOnWorkerReturns.result2, fake.findBuildContainerOnWorkerReturns.result3
}

func (fake *FakeTeam) FindBuildContainerOnWorkerCallCount() int {
	fake.findBuildContainerOnWorkerMutex.RLock()
	defer fake.findBuildContainerOnWorkerMutex.RUnlock()
	return len(fake.findBuildContainerOnWorkerArgsForCall)
}

func (fake *FakeTeam) FindBuildContainerOnWorkerArgsForCall(i int) (string, int, atc.PlanID) {
	fake.findBuildContainerOnWorkerMutex.RLock()
	defer fake.findBuildContainerOnWorkerMutex.RUnlock()
	return fake.findBuildContainerOnWorkerArgsForCall[i].workerName, fake.findBuildContainerOnWorkerArgsForCall[i].buildID, fake.findBuildContainerOnWorkerArgsForCall[i].planID
}

func (fake *FakeTeam) FindBuildContainerOnWorkerReturns(result1 dbng.CreatingContainer, result2 dbng.CreatedContainer, result3 error) {
	fake.FindBuildContainerOnWorkerStub = nil
	fake.findBuildContainerOnWorkerReturns = struct {
		result1 dbng.CreatingContainer
		result2 dbng.CreatedContainer
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeam) CreateBuildContainer(workerName string, buildID int, planID atc.PlanID, meta dbng.ContainerMetadata) (dbng.CreatingContainer, error) {
	fake.createBuildContainerMutex.Lock()
	fake.createBuildContainerArgsForCall = append(fake.createBuildContainerArgsForCall, struct {
		workerName string
		buildID    int
		planID     atc.PlanID
		meta       dbng.ContainerMetadata
	}{workerName, buildID, planID, meta})
	fake.recordInvocation("CreateBuildContainer", []interface{}{workerName, buildID, planID, meta})
	fake.createBuildContainerMutex.Unlock()
	if fake.CreateBuildContainerStub != nil {
		return fake.CreateBuildContainerStub(workerName, buildID, planID, meta)
	}
	return fake.createBuildContainerReturns.result1, fake.createBuildContainerReturns.result2
}

func (fake *FakeTeam) CreateBuildContainerCallCount() int {
	fake.createBuildContainerMutex.RLock()
	defer fake.createBuildContainerMutex.RUnlock()
	return len(fake.createBuildContainerArgsForCall)
}

func (fake *FakeTeam) CreateBuildContainerArgsForCall(i int) (string, int, atc.PlanID, dbng.ContainerMetadata) {
	fake.createBuildContainerMutex.RLock()
	defer fake.createBuildContainerMutex.RUnlock()
	return fake.createBuildContainerArgsForCall[i].workerName, fake.createBuildContainerArgsForCall[i].buildID, fake.createBuildContainerArgsForCall[i].planID, fake.createBuildContainerArgsForCall[i].meta
}

func (fake *FakeTeam) CreateBuildContainerReturns(result1 dbng.CreatingContainer, result2 error) {
	fake.CreateBuildContainerStub = nil
	fake.createBuildContainerReturns = struct {
		result1 dbng.CreatingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeTeam) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.savePipelineMutex.RLock()
	defer fake.savePipelineMutex.RUnlock()
	fake.findPipelineByNameMutex.RLock()
	defer fake.findPipelineByNameMutex.RUnlock()
	fake.createOneOffBuildMutex.RLock()
	defer fake.createOneOffBuildMutex.RUnlock()
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	fake.findContainerByHandleMutex.RLock()
	defer fake.findContainerByHandleMutex.RUnlock()
	fake.findWorkerForResourceCheckContainerMutex.RLock()
	defer fake.findWorkerForResourceCheckContainerMutex.RUnlock()
	fake.findResourceCheckContainerOnWorkerMutex.RLock()
	defer fake.findResourceCheckContainerOnWorkerMutex.RUnlock()
	fake.createResourceCheckContainerMutex.RLock()
	defer fake.createResourceCheckContainerMutex.RUnlock()
	fake.createResourceGetContainerMutex.RLock()
	defer fake.createResourceGetContainerMutex.RUnlock()
	fake.findWorkerForBuildContainerMutex.RLock()
	defer fake.findWorkerForBuildContainerMutex.RUnlock()
	fake.findBuildContainerOnWorkerMutex.RLock()
	defer fake.findBuildContainerOnWorkerMutex.RUnlock()
	fake.createBuildContainerMutex.RLock()
	defer fake.createBuildContainerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeam) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dbng.Team = new(FakeTeam)
