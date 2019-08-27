// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeCreatedContainer struct {
	DestroyingStub        func() (db.DestroyingContainer, error)
	destroyingMutex       sync.RWMutex
	destroyingArgsForCall []struct {
	}
	destroyingReturns struct {
		result1 db.DestroyingContainer
		result2 error
	}
	destroyingReturnsOnCall map[int]struct {
		result1 db.DestroyingContainer
		result2 error
	}
	DiscontinueStub        func() (db.DestroyingContainer, error)
	discontinueMutex       sync.RWMutex
	discontinueArgsForCall []struct {
	}
	discontinueReturns struct {
		result1 db.DestroyingContainer
		result2 error
	}
	discontinueReturnsOnCall map[int]struct {
		result1 db.DestroyingContainer
		result2 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	IsHijackedStub        func() bool
	isHijackedMutex       sync.RWMutex
	isHijackedArgsForCall []struct {
	}
	isHijackedReturns struct {
		result1 bool
	}
	isHijackedReturnsOnCall map[int]struct {
		result1 bool
	}
	MarkAsHijackedStub        func() error
	markAsHijackedMutex       sync.RWMutex
	markAsHijackedArgsForCall []struct {
	}
	markAsHijackedReturns struct {
		result1 error
	}
	markAsHijackedReturnsOnCall map[int]struct {
		result1 error
	}
	MetadataStub        func() db.ContainerMetadata
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct {
	}
	metadataReturns struct {
		result1 db.ContainerMetadata
	}
	metadataReturnsOnCall map[int]struct {
		result1 db.ContainerMetadata
	}
	StateStub        func() string
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
	}
	stateReturns struct {
		result1 string
	}
	stateReturnsOnCall map[int]struct {
		result1 string
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct {
	}
	teamNameReturns struct {
		result1 string
	}
	teamNameReturnsOnCall map[int]struct {
		result1 string
	}
	WorkerNameStub        func() string
	workerNameMutex       sync.RWMutex
	workerNameArgsForCall []struct {
	}
	workerNameReturns struct {
		result1 string
	}
	workerNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCreatedContainer) Destroying() (db.DestroyingContainer, error) {
	fake.destroyingMutex.Lock()
	ret, specificReturn := fake.destroyingReturnsOnCall[len(fake.destroyingArgsForCall)]
	fake.destroyingArgsForCall = append(fake.destroyingArgsForCall, struct {
	}{})
	fake.recordInvocation("Destroying", []interface{}{})
	fake.destroyingMutex.Unlock()
	if fake.DestroyingStub != nil {
		return fake.DestroyingStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.destroyingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreatedContainer) DestroyingCallCount() int {
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	return len(fake.destroyingArgsForCall)
}

func (fake *FakeCreatedContainer) DestroyingCalls(stub func() (db.DestroyingContainer, error)) {
	fake.destroyingMutex.Lock()
	defer fake.destroyingMutex.Unlock()
	fake.DestroyingStub = stub
}

func (fake *FakeCreatedContainer) DestroyingReturns(result1 db.DestroyingContainer, result2 error) {
	fake.destroyingMutex.Lock()
	defer fake.destroyingMutex.Unlock()
	fake.DestroyingStub = nil
	fake.destroyingReturns = struct {
		result1 db.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) DestroyingReturnsOnCall(i int, result1 db.DestroyingContainer, result2 error) {
	fake.destroyingMutex.Lock()
	defer fake.destroyingMutex.Unlock()
	fake.DestroyingStub = nil
	if fake.destroyingReturnsOnCall == nil {
		fake.destroyingReturnsOnCall = make(map[int]struct {
			result1 db.DestroyingContainer
			result2 error
		})
	}
	fake.destroyingReturnsOnCall[i] = struct {
		result1 db.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) Discontinue() (db.DestroyingContainer, error) {
	fake.discontinueMutex.Lock()
	ret, specificReturn := fake.discontinueReturnsOnCall[len(fake.discontinueArgsForCall)]
	fake.discontinueArgsForCall = append(fake.discontinueArgsForCall, struct {
	}{})
	fake.recordInvocation("Discontinue", []interface{}{})
	fake.discontinueMutex.Unlock()
	if fake.DiscontinueStub != nil {
		return fake.DiscontinueStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.discontinueReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCreatedContainer) DiscontinueCallCount() int {
	fake.discontinueMutex.RLock()
	defer fake.discontinueMutex.RUnlock()
	return len(fake.discontinueArgsForCall)
}

func (fake *FakeCreatedContainer) DiscontinueCalls(stub func() (db.DestroyingContainer, error)) {
	fake.discontinueMutex.Lock()
	defer fake.discontinueMutex.Unlock()
	fake.DiscontinueStub = stub
}

func (fake *FakeCreatedContainer) DiscontinueReturns(result1 db.DestroyingContainer, result2 error) {
	fake.discontinueMutex.Lock()
	defer fake.discontinueMutex.Unlock()
	fake.DiscontinueStub = nil
	fake.discontinueReturns = struct {
		result1 db.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) DiscontinueReturnsOnCall(i int, result1 db.DestroyingContainer, result2 error) {
	fake.discontinueMutex.Lock()
	defer fake.discontinueMutex.Unlock()
	fake.DiscontinueStub = nil
	if fake.discontinueReturnsOnCall == nil {
		fake.discontinueReturnsOnCall = make(map[int]struct {
			result1 db.DestroyingContainer
			result2 error
		})
	}
	fake.discontinueReturnsOnCall[i] = struct {
		result1 db.DestroyingContainer
		result2 error
	}{result1, result2}
}

func (fake *FakeCreatedContainer) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeCreatedContainer) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeCreatedContainer) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeCreatedContainer) IDCalls(stub func() int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeCreatedContainer) IDReturns(result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedContainer) IDReturnsOnCall(i int, result1 int) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCreatedContainer) IsHijacked() bool {
	fake.isHijackedMutex.Lock()
	ret, specificReturn := fake.isHijackedReturnsOnCall[len(fake.isHijackedArgsForCall)]
	fake.isHijackedArgsForCall = append(fake.isHijackedArgsForCall, struct {
	}{})
	fake.recordInvocation("IsHijacked", []interface{}{})
	fake.isHijackedMutex.Unlock()
	if fake.IsHijackedStub != nil {
		return fake.IsHijackedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isHijackedReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) IsHijackedCallCount() int {
	fake.isHijackedMutex.RLock()
	defer fake.isHijackedMutex.RUnlock()
	return len(fake.isHijackedArgsForCall)
}

func (fake *FakeCreatedContainer) IsHijackedCalls(stub func() bool) {
	fake.isHijackedMutex.Lock()
	defer fake.isHijackedMutex.Unlock()
	fake.IsHijackedStub = stub
}

func (fake *FakeCreatedContainer) IsHijackedReturns(result1 bool) {
	fake.isHijackedMutex.Lock()
	defer fake.isHijackedMutex.Unlock()
	fake.IsHijackedStub = nil
	fake.isHijackedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCreatedContainer) IsHijackedReturnsOnCall(i int, result1 bool) {
	fake.isHijackedMutex.Lock()
	defer fake.isHijackedMutex.Unlock()
	fake.IsHijackedStub = nil
	if fake.isHijackedReturnsOnCall == nil {
		fake.isHijackedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isHijackedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCreatedContainer) MarkAsHijacked() error {
	fake.markAsHijackedMutex.Lock()
	ret, specificReturn := fake.markAsHijackedReturnsOnCall[len(fake.markAsHijackedArgsForCall)]
	fake.markAsHijackedArgsForCall = append(fake.markAsHijackedArgsForCall, struct {
	}{})
	fake.recordInvocation("MarkAsHijacked", []interface{}{})
	fake.markAsHijackedMutex.Unlock()
	if fake.MarkAsHijackedStub != nil {
		return fake.MarkAsHijackedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.markAsHijackedReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) MarkAsHijackedCallCount() int {
	fake.markAsHijackedMutex.RLock()
	defer fake.markAsHijackedMutex.RUnlock()
	return len(fake.markAsHijackedArgsForCall)
}

func (fake *FakeCreatedContainer) MarkAsHijackedCalls(stub func() error) {
	fake.markAsHijackedMutex.Lock()
	defer fake.markAsHijackedMutex.Unlock()
	fake.MarkAsHijackedStub = stub
}

func (fake *FakeCreatedContainer) MarkAsHijackedReturns(result1 error) {
	fake.markAsHijackedMutex.Lock()
	defer fake.markAsHijackedMutex.Unlock()
	fake.MarkAsHijackedStub = nil
	fake.markAsHijackedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedContainer) MarkAsHijackedReturnsOnCall(i int, result1 error) {
	fake.markAsHijackedMutex.Lock()
	defer fake.markAsHijackedMutex.Unlock()
	fake.MarkAsHijackedStub = nil
	if fake.markAsHijackedReturnsOnCall == nil {
		fake.markAsHijackedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.markAsHijackedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCreatedContainer) Metadata() db.ContainerMetadata {
	fake.metadataMutex.Lock()
	ret, specificReturn := fake.metadataReturnsOnCall[len(fake.metadataArgsForCall)]
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct {
	}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metadataReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeCreatedContainer) MetadataCalls(stub func() db.ContainerMetadata) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = stub
}

func (fake *FakeCreatedContainer) MetadataReturns(result1 db.ContainerMetadata) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 db.ContainerMetadata
	}{result1}
}

func (fake *FakeCreatedContainer) MetadataReturnsOnCall(i int, result1 db.ContainerMetadata) {
	fake.metadataMutex.Lock()
	defer fake.metadataMutex.Unlock()
	fake.MetadataStub = nil
	if fake.metadataReturnsOnCall == nil {
		fake.metadataReturnsOnCall = make(map[int]struct {
			result1 db.ContainerMetadata
		})
	}
	fake.metadataReturnsOnCall[i] = struct {
		result1 db.ContainerMetadata
	}{result1}
}

func (fake *FakeCreatedContainer) State() string {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
	}{})
	fake.recordInvocation("State", []interface{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stateReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeCreatedContainer) StateCalls(stub func() string) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = stub
}

func (fake *FakeCreatedContainer) StateReturns(result1 string) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) StateReturnsOnCall(i int, result1 string) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) TeamName() string {
	fake.teamNameMutex.Lock()
	ret, specificReturn := fake.teamNameReturnsOnCall[len(fake.teamNameArgsForCall)]
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct {
	}{})
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if fake.TeamNameStub != nil {
		return fake.TeamNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.teamNameReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeCreatedContainer) TeamNameCalls(stub func() string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = stub
}

func (fake *FakeCreatedContainer) TeamNameReturns(result1 string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) TeamNameReturnsOnCall(i int, result1 string) {
	fake.teamNameMutex.Lock()
	defer fake.teamNameMutex.Unlock()
	fake.TeamNameStub = nil
	if fake.teamNameReturnsOnCall == nil {
		fake.teamNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.teamNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) WorkerName() string {
	fake.workerNameMutex.Lock()
	ret, specificReturn := fake.workerNameReturnsOnCall[len(fake.workerNameArgsForCall)]
	fake.workerNameArgsForCall = append(fake.workerNameArgsForCall, struct {
	}{})
	fake.recordInvocation("WorkerName", []interface{}{})
	fake.workerNameMutex.Unlock()
	if fake.WorkerNameStub != nil {
		return fake.WorkerNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.workerNameReturns
	return fakeReturns.result1
}

func (fake *FakeCreatedContainer) WorkerNameCallCount() int {
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	return len(fake.workerNameArgsForCall)
}

func (fake *FakeCreatedContainer) WorkerNameCalls(stub func() string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = stub
}

func (fake *FakeCreatedContainer) WorkerNameReturns(result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	fake.workerNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) WorkerNameReturnsOnCall(i int, result1 string) {
	fake.workerNameMutex.Lock()
	defer fake.workerNameMutex.Unlock()
	fake.WorkerNameStub = nil
	if fake.workerNameReturnsOnCall == nil {
		fake.workerNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.workerNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCreatedContainer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.destroyingMutex.RLock()
	defer fake.destroyingMutex.RUnlock()
	fake.discontinueMutex.RLock()
	defer fake.discontinueMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isHijackedMutex.RLock()
	defer fake.isHijackedMutex.RUnlock()
	fake.markAsHijackedMutex.RLock()
	defer fake.markAsHijackedMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.workerNameMutex.RLock()
	defer fake.workerNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCreatedContainer) recordInvocation(key string, args []interface{}) {
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

var _ db.CreatedContainer = new(FakeCreatedContainer)
