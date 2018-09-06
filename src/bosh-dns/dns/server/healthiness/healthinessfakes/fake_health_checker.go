// Code generated by counterfeiter. DO NOT EDIT.
package healthinessfakes

import (
	"bosh-dns/dns/server/healthiness"
	"sync"
)

type FakeHealthChecker struct {
	GetStatusStub        func(ip string) healthiness.HealthState
	getStatusMutex       sync.RWMutex
	getStatusArgsForCall []struct {
		ip string
	}
	getStatusReturns struct {
		result1 healthiness.HealthState
	}
	getStatusReturnsOnCall map[int]struct {
		result1 healthiness.HealthState
	}
	GetStatus2Stub        func(ip string) string
	getStatus2Mutex       sync.RWMutex
	getStatus2ArgsForCall []struct {
		ip string
	}
	getStatus2Returns struct {
		result1 string
	}
	getStatus2ReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHealthChecker) GetStatus(ip string) healthiness.HealthState {
	fake.getStatusMutex.Lock()
	ret, specificReturn := fake.getStatusReturnsOnCall[len(fake.getStatusArgsForCall)]
	fake.getStatusArgsForCall = append(fake.getStatusArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("GetStatus", []interface{}{ip})
	fake.getStatusMutex.Unlock()
	if fake.GetStatusStub != nil {
		return fake.GetStatusStub(ip)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getStatusReturns.result1
}

func (fake *FakeHealthChecker) GetStatusCallCount() int {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return len(fake.getStatusArgsForCall)
}

func (fake *FakeHealthChecker) GetStatusArgsForCall(i int) string {
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	return fake.getStatusArgsForCall[i].ip
}

func (fake *FakeHealthChecker) GetStatusReturns(result1 healthiness.HealthState) {
	fake.GetStatusStub = nil
	fake.getStatusReturns = struct {
		result1 healthiness.HealthState
	}{result1}
}

func (fake *FakeHealthChecker) GetStatusReturnsOnCall(i int, result1 healthiness.HealthState) {
	fake.GetStatusStub = nil
	if fake.getStatusReturnsOnCall == nil {
		fake.getStatusReturnsOnCall = make(map[int]struct {
			result1 healthiness.HealthState
		})
	}
	fake.getStatusReturnsOnCall[i] = struct {
		result1 healthiness.HealthState
	}{result1}
}

func (fake *FakeHealthChecker) GetStatus2(ip string) string {
	fake.getStatus2Mutex.Lock()
	ret, specificReturn := fake.getStatus2ReturnsOnCall[len(fake.getStatus2ArgsForCall)]
	fake.getStatus2ArgsForCall = append(fake.getStatus2ArgsForCall, struct {
		ip string
	}{ip})
	fake.recordInvocation("GetStatus2", []interface{}{ip})
	fake.getStatus2Mutex.Unlock()
	if fake.GetStatus2Stub != nil {
		return fake.GetStatus2Stub(ip)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getStatus2Returns.result1
}

func (fake *FakeHealthChecker) GetStatus2CallCount() int {
	fake.getStatus2Mutex.RLock()
	defer fake.getStatus2Mutex.RUnlock()
	return len(fake.getStatus2ArgsForCall)
}

func (fake *FakeHealthChecker) GetStatus2ArgsForCall(i int) string {
	fake.getStatus2Mutex.RLock()
	defer fake.getStatus2Mutex.RUnlock()
	return fake.getStatus2ArgsForCall[i].ip
}

func (fake *FakeHealthChecker) GetStatus2Returns(result1 string) {
	fake.GetStatus2Stub = nil
	fake.getStatus2Returns = struct {
		result1 string
	}{result1}
}

func (fake *FakeHealthChecker) GetStatus2ReturnsOnCall(i int, result1 string) {
	fake.GetStatus2Stub = nil
	if fake.getStatus2ReturnsOnCall == nil {
		fake.getStatus2ReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getStatus2ReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeHealthChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStatusMutex.RLock()
	defer fake.getStatusMutex.RUnlock()
	fake.getStatus2Mutex.RLock()
	defer fake.getStatus2Mutex.RUnlock()
	return fake.invocations
}

func (fake *FakeHealthChecker) recordInvocation(key string, args []interface{}) {
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

var _ healthiness.HealthChecker = new(FakeHealthChecker)
