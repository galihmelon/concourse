// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

type FakeCheckFactory struct {
	ChecksStub        func() ([]db.Check, error)
	checksMutex       sync.RWMutex
	checksArgsForCall []struct {
	}
	checksReturns struct {
		result1 []db.Check
		result2 error
	}
	checksReturnsOnCall map[int]struct {
		result1 []db.Check
		result2 error
	}
	CreateCheckStub        func(int, string) (db.Check, error)
	createCheckMutex       sync.RWMutex
	createCheckArgsForCall []struct {
		arg1 int
		arg2 string
	}
	createCheckReturns struct {
		result1 db.Check
		result2 error
	}
	createCheckReturnsOnCall map[int]struct {
		result1 db.Check
		result2 error
	}
	CreateCheckFromVersionStub        func(int, string, atc.Version) (db.Check, error)
	createCheckFromVersionMutex       sync.RWMutex
	createCheckFromVersionArgsForCall []struct {
		arg1 int
		arg2 string
		arg3 atc.Version
	}
	createCheckFromVersionReturns struct {
		result1 db.Check
		result2 error
	}
	createCheckFromVersionReturnsOnCall map[int]struct {
		result1 db.Check
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCheckFactory) Checks() ([]db.Check, error) {
	fake.checksMutex.Lock()
	ret, specificReturn := fake.checksReturnsOnCall[len(fake.checksArgsForCall)]
	fake.checksArgsForCall = append(fake.checksArgsForCall, struct {
	}{})
	fake.recordInvocation("Checks", []interface{}{})
	fake.checksMutex.Unlock()
	if fake.ChecksStub != nil {
		return fake.ChecksStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.checksReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) ChecksCallCount() int {
	fake.checksMutex.RLock()
	defer fake.checksMutex.RUnlock()
	return len(fake.checksArgsForCall)
}

func (fake *FakeCheckFactory) ChecksCalls(stub func() ([]db.Check, error)) {
	fake.checksMutex.Lock()
	defer fake.checksMutex.Unlock()
	fake.ChecksStub = stub
}

func (fake *FakeCheckFactory) ChecksReturns(result1 []db.Check, result2 error) {
	fake.checksMutex.Lock()
	defer fake.checksMutex.Unlock()
	fake.ChecksStub = nil
	fake.checksReturns = struct {
		result1 []db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) ChecksReturnsOnCall(i int, result1 []db.Check, result2 error) {
	fake.checksMutex.Lock()
	defer fake.checksMutex.Unlock()
	fake.ChecksStub = nil
	if fake.checksReturnsOnCall == nil {
		fake.checksReturnsOnCall = make(map[int]struct {
			result1 []db.Check
			result2 error
		})
	}
	fake.checksReturnsOnCall[i] = struct {
		result1 []db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) CreateCheck(arg1 int, arg2 string) (db.Check, error) {
	fake.createCheckMutex.Lock()
	ret, specificReturn := fake.createCheckReturnsOnCall[len(fake.createCheckArgsForCall)]
	fake.createCheckArgsForCall = append(fake.createCheckArgsForCall, struct {
		arg1 int
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateCheck", []interface{}{arg1, arg2})
	fake.createCheckMutex.Unlock()
	if fake.CreateCheckStub != nil {
		return fake.CreateCheckStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createCheckReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) CreateCheckCallCount() int {
	fake.createCheckMutex.RLock()
	defer fake.createCheckMutex.RUnlock()
	return len(fake.createCheckArgsForCall)
}

func (fake *FakeCheckFactory) CreateCheckCalls(stub func(int, string) (db.Check, error)) {
	fake.createCheckMutex.Lock()
	defer fake.createCheckMutex.Unlock()
	fake.CreateCheckStub = stub
}

func (fake *FakeCheckFactory) CreateCheckArgsForCall(i int) (int, string) {
	fake.createCheckMutex.RLock()
	defer fake.createCheckMutex.RUnlock()
	argsForCall := fake.createCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCheckFactory) CreateCheckReturns(result1 db.Check, result2 error) {
	fake.createCheckMutex.Lock()
	defer fake.createCheckMutex.Unlock()
	fake.CreateCheckStub = nil
	fake.createCheckReturns = struct {
		result1 db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) CreateCheckReturnsOnCall(i int, result1 db.Check, result2 error) {
	fake.createCheckMutex.Lock()
	defer fake.createCheckMutex.Unlock()
	fake.CreateCheckStub = nil
	if fake.createCheckReturnsOnCall == nil {
		fake.createCheckReturnsOnCall = make(map[int]struct {
			result1 db.Check
			result2 error
		})
	}
	fake.createCheckReturnsOnCall[i] = struct {
		result1 db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) CreateCheckFromVersion(arg1 int, arg2 string, arg3 atc.Version) (db.Check, error) {
	fake.createCheckFromVersionMutex.Lock()
	ret, specificReturn := fake.createCheckFromVersionReturnsOnCall[len(fake.createCheckFromVersionArgsForCall)]
	fake.createCheckFromVersionArgsForCall = append(fake.createCheckFromVersionArgsForCall, struct {
		arg1 int
		arg2 string
		arg3 atc.Version
	}{arg1, arg2, arg3})
	fake.recordInvocation("CreateCheckFromVersion", []interface{}{arg1, arg2, arg3})
	fake.createCheckFromVersionMutex.Unlock()
	if fake.CreateCheckFromVersionStub != nil {
		return fake.CreateCheckFromVersionStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createCheckFromVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCheckFactory) CreateCheckFromVersionCallCount() int {
	fake.createCheckFromVersionMutex.RLock()
	defer fake.createCheckFromVersionMutex.RUnlock()
	return len(fake.createCheckFromVersionArgsForCall)
}

func (fake *FakeCheckFactory) CreateCheckFromVersionCalls(stub func(int, string, atc.Version) (db.Check, error)) {
	fake.createCheckFromVersionMutex.Lock()
	defer fake.createCheckFromVersionMutex.Unlock()
	fake.CreateCheckFromVersionStub = stub
}

func (fake *FakeCheckFactory) CreateCheckFromVersionArgsForCall(i int) (int, string, atc.Version) {
	fake.createCheckFromVersionMutex.RLock()
	defer fake.createCheckFromVersionMutex.RUnlock()
	argsForCall := fake.createCheckFromVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCheckFactory) CreateCheckFromVersionReturns(result1 db.Check, result2 error) {
	fake.createCheckFromVersionMutex.Lock()
	defer fake.createCheckFromVersionMutex.Unlock()
	fake.CreateCheckFromVersionStub = nil
	fake.createCheckFromVersionReturns = struct {
		result1 db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) CreateCheckFromVersionReturnsOnCall(i int, result1 db.Check, result2 error) {
	fake.createCheckFromVersionMutex.Lock()
	defer fake.createCheckFromVersionMutex.Unlock()
	fake.CreateCheckFromVersionStub = nil
	if fake.createCheckFromVersionReturnsOnCall == nil {
		fake.createCheckFromVersionReturnsOnCall = make(map[int]struct {
			result1 db.Check
			result2 error
		})
	}
	fake.createCheckFromVersionReturnsOnCall[i] = struct {
		result1 db.Check
		result2 error
	}{result1, result2}
}

func (fake *FakeCheckFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checksMutex.RLock()
	defer fake.checksMutex.RUnlock()
	fake.createCheckMutex.RLock()
	defer fake.createCheckMutex.RUnlock()
	fake.createCheckFromVersionMutex.RLock()
	defer fake.createCheckFromVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCheckFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.CheckFactory = new(FakeCheckFactory)
