// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeReleaseClient struct {
	ListStub        func(productSlug string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		productSlug string
	}
	listReturns struct {
		result1 error
	}
	GetStub        func(productSlug string, releaseVersion string) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	getReturns struct {
		result1 error
	}
	CreateStub        func(productSlug string, releaseVersion string, releaseType string, eulaSlug string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		productSlug    string
		releaseVersion string
		releaseType    string
		eulaSlug       string
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(productSlug string, releaseVersion string, availability *string) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		productSlug    string
		releaseVersion string
		availability   *string
	}
	updateReturns struct {
		result1 error
	}
	DeleteStub        func(productSlug string, releaseVersion string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseClient) List(productSlug string) error {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		productSlug string
	}{productSlug})
	fake.recordInvocation("List", []interface{}{productSlug})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(productSlug)
	} else {
		return fake.listReturns.result1
	}
}

func (fake *FakeReleaseClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeReleaseClient) ListArgsForCall(i int) string {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug
}

func (fake *FakeReleaseClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseClient) Get(productSlug string, releaseVersion string) error {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("Get", []interface{}{productSlug, releaseVersion})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(productSlug, releaseVersion)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeReleaseClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeReleaseClient) GetArgsForCall(i int) (string, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].productSlug, fake.getArgsForCall[i].releaseVersion
}

func (fake *FakeReleaseClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseClient) Create(productSlug string, releaseVersion string, releaseType string, eulaSlug string) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		productSlug    string
		releaseVersion string
		releaseType    string
		eulaSlug       string
	}{productSlug, releaseVersion, releaseType, eulaSlug})
	fake.recordInvocation("Create", []interface{}{productSlug, releaseVersion, releaseType, eulaSlug})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(productSlug, releaseVersion, releaseType, eulaSlug)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeReleaseClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeReleaseClient) CreateArgsForCall(i int) (string, string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].productSlug, fake.createArgsForCall[i].releaseVersion, fake.createArgsForCall[i].releaseType, fake.createArgsForCall[i].eulaSlug
}

func (fake *FakeReleaseClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseClient) Update(productSlug string, releaseVersion string, availability *string) error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		productSlug    string
		releaseVersion string
		availability   *string
	}{productSlug, releaseVersion, availability})
	fake.recordInvocation("Update", []interface{}{productSlug, releaseVersion, availability})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(productSlug, releaseVersion, availability)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeReleaseClient) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeReleaseClient) UpdateArgsForCall(i int) (string, string, *string) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].productSlug, fake.updateArgsForCall[i].releaseVersion, fake.updateArgsForCall[i].availability
}

func (fake *FakeReleaseClient) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseClient) Delete(productSlug string, releaseVersion string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("Delete", []interface{}{productSlug, releaseVersion})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(productSlug, releaseVersion)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeReleaseClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeReleaseClient) DeleteArgsForCall(i int) (string, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].productSlug, fake.deleteArgsForCall[i].releaseVersion
}

func (fake *FakeReleaseClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeReleaseClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.ReleaseClient = new(FakeReleaseClient)
