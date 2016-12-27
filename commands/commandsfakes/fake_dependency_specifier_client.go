// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
)

type FakeDependencySpecifierClient struct {
	ListStub        func(productSlug string, releaseVersion string) error
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	listReturns struct {
		result1 error
	}
	GetStub        func(productSlug string, releaseVersion string, dependencySpecifierID int) error
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		productSlug           string
		releaseVersion        string
		dependencySpecifierID int
	}
	getReturns struct {
		result1 error
	}
	CreateStub        func(productSlug string, releaseVersion string, dependentProductSlug string, specifier string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		productSlug          string
		releaseVersion       string
		dependentProductSlug string
		specifier            string
	}
	createReturns struct {
		result1 error
	}
	DeleteStub        func(productSlug string, releaseVersion string, dependencyspecifierID int) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		productSlug           string
		releaseVersion        string
		dependencyspecifierID int
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDependencySpecifierClient) List(productSlug string, releaseVersion string) error {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("List", []interface{}{productSlug, releaseVersion})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(productSlug, releaseVersion)
	} else {
		return fake.listReturns.result1
	}
}

func (fake *FakeDependencySpecifierClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeDependencySpecifierClient) ListArgsForCall(i int) (string, string) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].productSlug, fake.listArgsForCall[i].releaseVersion
}

func (fake *FakeDependencySpecifierClient) ListReturns(result1 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencySpecifierClient) Get(productSlug string, releaseVersion string, dependencySpecifierID int) error {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		productSlug           string
		releaseVersion        string
		dependencySpecifierID int
	}{productSlug, releaseVersion, dependencySpecifierID})
	fake.recordInvocation("Get", []interface{}{productSlug, releaseVersion, dependencySpecifierID})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(productSlug, releaseVersion, dependencySpecifierID)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeDependencySpecifierClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeDependencySpecifierClient) GetArgsForCall(i int) (string, string, int) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].productSlug, fake.getArgsForCall[i].releaseVersion, fake.getArgsForCall[i].dependencySpecifierID
}

func (fake *FakeDependencySpecifierClient) GetReturns(result1 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencySpecifierClient) Create(productSlug string, releaseVersion string, dependentProductSlug string, specifier string) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		productSlug          string
		releaseVersion       string
		dependentProductSlug string
		specifier            string
	}{productSlug, releaseVersion, dependentProductSlug, specifier})
	fake.recordInvocation("Create", []interface{}{productSlug, releaseVersion, dependentProductSlug, specifier})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(productSlug, releaseVersion, dependentProductSlug, specifier)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeDependencySpecifierClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeDependencySpecifierClient) CreateArgsForCall(i int) (string, string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].productSlug, fake.createArgsForCall[i].releaseVersion, fake.createArgsForCall[i].dependentProductSlug, fake.createArgsForCall[i].specifier
}

func (fake *FakeDependencySpecifierClient) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencySpecifierClient) Delete(productSlug string, releaseVersion string, dependencyspecifierID int) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		productSlug           string
		releaseVersion        string
		dependencyspecifierID int
	}{productSlug, releaseVersion, dependencyspecifierID})
	fake.recordInvocation("Delete", []interface{}{productSlug, releaseVersion, dependencyspecifierID})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(productSlug, releaseVersion, dependencyspecifierID)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeDependencySpecifierClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeDependencySpecifierClient) DeleteArgsForCall(i int) (string, string, int) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].productSlug, fake.deleteArgsForCall[i].releaseVersion, fake.deleteArgsForCall[i].dependencyspecifierID
}

func (fake *FakeDependencySpecifierClient) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDependencySpecifierClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDependencySpecifierClient) recordInvocation(key string, args []interface{}) {
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

var _ commands.DependencySpecifierClient = new(FakeDependencySpecifierClient)
