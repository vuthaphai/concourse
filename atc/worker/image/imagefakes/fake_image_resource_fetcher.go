// Code generated by counterfeiter. DO NOT EDIT.
package imagefakes

import (
	context "context"
	io "io"
	sync "sync"

	lager "code.cloudfoundry.org/lager"
	atc "github.com/concourse/concourse/atc"
	db "github.com/concourse/concourse/atc/db"
	worker "github.com/concourse/concourse/atc/worker"
	image "github.com/concourse/concourse/atc/worker/image"
)

type FakeImageResourceFetcher struct {
	FetchStub        func(context.Context, lager.Logger, db.CreatingContainer, bool) (worker.Volume, io.ReadCloser, atc.Version, error)
	fetchMutex       sync.RWMutex
	fetchArgsForCall []struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.CreatingContainer
		arg4 bool
	}
	fetchReturns struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}
	fetchReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageResourceFetcher) Fetch(arg1 context.Context, arg2 lager.Logger, arg3 db.CreatingContainer, arg4 bool) (worker.Volume, io.ReadCloser, atc.Version, error) {
	fake.fetchMutex.Lock()
	ret, specificReturn := fake.fetchReturnsOnCall[len(fake.fetchArgsForCall)]
	fake.fetchArgsForCall = append(fake.fetchArgsForCall, struct {
		arg1 context.Context
		arg2 lager.Logger
		arg3 db.CreatingContainer
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Fetch", []interface{}{arg1, arg2, arg3, arg4})
	fake.fetchMutex.Unlock()
	if fake.FetchStub != nil {
		return fake.FetchStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.fetchReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeImageResourceFetcher) FetchCallCount() int {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	return len(fake.fetchArgsForCall)
}

func (fake *FakeImageResourceFetcher) FetchArgsForCall(i int) (context.Context, lager.Logger, db.CreatingContainer, bool) {
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	argsForCall := fake.fetchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeImageResourceFetcher) FetchReturns(result1 worker.Volume, result2 io.ReadCloser, result3 atc.Version, result4 error) {
	fake.FetchStub = nil
	fake.fetchReturns = struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeImageResourceFetcher) FetchReturnsOnCall(i int, result1 worker.Volume, result2 io.ReadCloser, result3 atc.Version, result4 error) {
	fake.FetchStub = nil
	if fake.fetchReturnsOnCall == nil {
		fake.fetchReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 io.ReadCloser
			result3 atc.Version
			result4 error
		})
	}
	fake.fetchReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 io.ReadCloser
		result3 atc.Version
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeImageResourceFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchMutex.RLock()
	defer fake.fetchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageResourceFetcher) recordInvocation(key string, args []interface{}) {
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

var _ image.ImageResourceFetcher = new(FakeImageResourceFetcher)
