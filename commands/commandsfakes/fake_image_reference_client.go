// Code generated by counterfeiter. DO NOT EDIT.
package commandsfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v2"
	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeImageReferenceClient struct {
	CreateStub        func(config pivnet.CreateImageReferenceConfig) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		config pivnet.CreateImageReferenceConfig
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageReferenceClient) Create(config pivnet.CreateImageReferenceConfig) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		config pivnet.CreateImageReferenceConfig
	}{config})
	fake.recordInvocation("Create", []interface{}{config})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(config)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeImageReferenceClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeImageReferenceClient) CreateArgsForCall(i int) pivnet.CreateImageReferenceConfig {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].config
}

func (fake *FakeImageReferenceClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageReferenceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageReferenceClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.ImageReferenceClient = new(FakeImageReferenceClient)