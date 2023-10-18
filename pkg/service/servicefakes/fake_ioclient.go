// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"context"
	"sync"

	"github.com/livekit/livekit-server/pkg/service"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/rpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FakeIOClient struct {
	CreateEgressStub        func(context.Context, *livekit.EgressInfo) (*emptypb.Empty, error)
	createEgressMutex       sync.RWMutex
	createEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}
	createEgressReturns struct {
		result1 *emptypb.Empty
		result2 error
	}
	createEgressReturnsOnCall map[int]struct {
		result1 *emptypb.Empty
		result2 error
	}
	GetEgressStub        func(context.Context, *rpc.GetEgressRequest) (*livekit.EgressInfo, error)
	getEgressMutex       sync.RWMutex
	getEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *rpc.GetEgressRequest
	}
	getEgressReturns struct {
		result1 *livekit.EgressInfo
		result2 error
	}
	getEgressReturnsOnCall map[int]struct {
		result1 *livekit.EgressInfo
		result2 error
	}
	ListEgressStub        func(context.Context, *livekit.ListEgressRequest) (*livekit.ListEgressResponse, error)
	listEgressMutex       sync.RWMutex
	listEgressArgsForCall []struct {
		arg1 context.Context
		arg2 *livekit.ListEgressRequest
	}
	listEgressReturns struct {
		result1 *livekit.ListEgressResponse
		result2 error
	}
	listEgressReturnsOnCall map[int]struct {
		result1 *livekit.ListEgressResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIOClient) CreateEgress(arg1 context.Context, arg2 *livekit.EgressInfo) (*emptypb.Empty, error) {
	fake.createEgressMutex.Lock()
	ret, specificReturn := fake.createEgressReturnsOnCall[len(fake.createEgressArgsForCall)]
	fake.createEgressArgsForCall = append(fake.createEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.EgressInfo
	}{arg1, arg2})
	stub := fake.CreateEgressStub
	fakeReturns := fake.createEgressReturns
	fake.recordInvocation("CreateEgress", []interface{}{arg1, arg2})
	fake.createEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIOClient) CreateEgressCallCount() int {
	fake.createEgressMutex.RLock()
	defer fake.createEgressMutex.RUnlock()
	return len(fake.createEgressArgsForCall)
}

func (fake *FakeIOClient) CreateEgressCalls(stub func(context.Context, *livekit.EgressInfo) (*emptypb.Empty, error)) {
	fake.createEgressMutex.Lock()
	defer fake.createEgressMutex.Unlock()
	fake.CreateEgressStub = stub
}

func (fake *FakeIOClient) CreateEgressArgsForCall(i int) (context.Context, *livekit.EgressInfo) {
	fake.createEgressMutex.RLock()
	defer fake.createEgressMutex.RUnlock()
	argsForCall := fake.createEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIOClient) CreateEgressReturns(result1 *emptypb.Empty, result2 error) {
	fake.createEgressMutex.Lock()
	defer fake.createEgressMutex.Unlock()
	fake.CreateEgressStub = nil
	fake.createEgressReturns = struct {
		result1 *emptypb.Empty
		result2 error
	}{result1, result2}
}

func (fake *FakeIOClient) CreateEgressReturnsOnCall(i int, result1 *emptypb.Empty, result2 error) {
	fake.createEgressMutex.Lock()
	defer fake.createEgressMutex.Unlock()
	fake.CreateEgressStub = nil
	if fake.createEgressReturnsOnCall == nil {
		fake.createEgressReturnsOnCall = make(map[int]struct {
			result1 *emptypb.Empty
			result2 error
		})
	}
	fake.createEgressReturnsOnCall[i] = struct {
		result1 *emptypb.Empty
		result2 error
	}{result1, result2}
}

func (fake *FakeIOClient) GetEgress(arg1 context.Context, arg2 *rpc.GetEgressRequest) (*livekit.EgressInfo, error) {
	fake.getEgressMutex.Lock()
	ret, specificReturn := fake.getEgressReturnsOnCall[len(fake.getEgressArgsForCall)]
	fake.getEgressArgsForCall = append(fake.getEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *rpc.GetEgressRequest
	}{arg1, arg2})
	stub := fake.GetEgressStub
	fakeReturns := fake.getEgressReturns
	fake.recordInvocation("GetEgress", []interface{}{arg1, arg2})
	fake.getEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIOClient) GetEgressCallCount() int {
	fake.getEgressMutex.RLock()
	defer fake.getEgressMutex.RUnlock()
	return len(fake.getEgressArgsForCall)
}

func (fake *FakeIOClient) GetEgressCalls(stub func(context.Context, *rpc.GetEgressRequest) (*livekit.EgressInfo, error)) {
	fake.getEgressMutex.Lock()
	defer fake.getEgressMutex.Unlock()
	fake.GetEgressStub = stub
}

func (fake *FakeIOClient) GetEgressArgsForCall(i int) (context.Context, *rpc.GetEgressRequest) {
	fake.getEgressMutex.RLock()
	defer fake.getEgressMutex.RUnlock()
	argsForCall := fake.getEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIOClient) GetEgressReturns(result1 *livekit.EgressInfo, result2 error) {
	fake.getEgressMutex.Lock()
	defer fake.getEgressMutex.Unlock()
	fake.GetEgressStub = nil
	fake.getEgressReturns = struct {
		result1 *livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeIOClient) GetEgressReturnsOnCall(i int, result1 *livekit.EgressInfo, result2 error) {
	fake.getEgressMutex.Lock()
	defer fake.getEgressMutex.Unlock()
	fake.GetEgressStub = nil
	if fake.getEgressReturnsOnCall == nil {
		fake.getEgressReturnsOnCall = make(map[int]struct {
			result1 *livekit.EgressInfo
			result2 error
		})
	}
	fake.getEgressReturnsOnCall[i] = struct {
		result1 *livekit.EgressInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeIOClient) ListEgress(arg1 context.Context, arg2 *livekit.ListEgressRequest) (*livekit.ListEgressResponse, error) {
	fake.listEgressMutex.Lock()
	ret, specificReturn := fake.listEgressReturnsOnCall[len(fake.listEgressArgsForCall)]
	fake.listEgressArgsForCall = append(fake.listEgressArgsForCall, struct {
		arg1 context.Context
		arg2 *livekit.ListEgressRequest
	}{arg1, arg2})
	stub := fake.ListEgressStub
	fakeReturns := fake.listEgressReturns
	fake.recordInvocation("ListEgress", []interface{}{arg1, arg2})
	fake.listEgressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIOClient) ListEgressCallCount() int {
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	return len(fake.listEgressArgsForCall)
}

func (fake *FakeIOClient) ListEgressCalls(stub func(context.Context, *livekit.ListEgressRequest) (*livekit.ListEgressResponse, error)) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = stub
}

func (fake *FakeIOClient) ListEgressArgsForCall(i int) (context.Context, *livekit.ListEgressRequest) {
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	argsForCall := fake.listEgressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIOClient) ListEgressReturns(result1 *livekit.ListEgressResponse, result2 error) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = nil
	fake.listEgressReturns = struct {
		result1 *livekit.ListEgressResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeIOClient) ListEgressReturnsOnCall(i int, result1 *livekit.ListEgressResponse, result2 error) {
	fake.listEgressMutex.Lock()
	defer fake.listEgressMutex.Unlock()
	fake.ListEgressStub = nil
	if fake.listEgressReturnsOnCall == nil {
		fake.listEgressReturnsOnCall = make(map[int]struct {
			result1 *livekit.ListEgressResponse
			result2 error
		})
	}
	fake.listEgressReturnsOnCall[i] = struct {
		result1 *livekit.ListEgressResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeIOClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createEgressMutex.RLock()
	defer fake.createEgressMutex.RUnlock()
	fake.getEgressMutex.RLock()
	defer fake.getEgressMutex.RUnlock()
	fake.listEgressMutex.RLock()
	defer fake.listEgressMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIOClient) recordInvocation(key string, args []interface{}) {
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

var _ service.IOClient = new(FakeIOClient)
