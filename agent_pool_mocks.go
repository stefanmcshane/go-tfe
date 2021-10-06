// Code generated by MockGen. DO NOT EDIT.
// Source: agent_pool.go

// Package tfe is a generated GoMock package.
package tfe

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAgentPools is a mock of AgentPools interface.
type MockAgentPools struct {
	ctrl     *gomock.Controller
	recorder *MockAgentPoolsMockRecorder
}

// MockAgentPoolsMockRecorder is the mock recorder for MockAgentPools.
type MockAgentPoolsMockRecorder struct {
	mock *MockAgentPools
}

// NewMockAgentPools creates a new mock instance.
func NewMockAgentPools(ctrl *gomock.Controller) *MockAgentPools {
	mock := &MockAgentPools{ctrl: ctrl}
	mock.recorder = &MockAgentPoolsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentPools) EXPECT() *MockAgentPoolsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAgentPools) Create(ctx context.Context, organization string, options AgentPoolCreateOptions) (*AgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, organization, options)
	ret0, _ := ret[0].(*AgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAgentPoolsMockRecorder) Create(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAgentPools)(nil).Create), ctx, organization, options)
}

// Delete mocks base method.
func (m *MockAgentPools) Delete(ctx context.Context, agentPoolID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, agentPoolID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAgentPoolsMockRecorder) Delete(ctx, agentPoolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAgentPools)(nil).Delete), ctx, agentPoolID)
}

// List mocks base method.
func (m *MockAgentPools) List(ctx context.Context, organization string, options AgentPoolListOptions) (*AgentPoolList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, organization, options)
	ret0, _ := ret[0].(*AgentPoolList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAgentPoolsMockRecorder) List(ctx, organization, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAgentPools)(nil).List), ctx, organization, options)
}

// Read mocks base method.
func (m *MockAgentPools) Read(ctx context.Context, agentPoolID string) (*AgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx, agentPoolID)
	ret0, _ := ret[0].(*AgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockAgentPoolsMockRecorder) Read(ctx, agentPoolID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockAgentPools)(nil).Read), ctx, agentPoolID)
}

// Update mocks base method.
func (m *MockAgentPools) Update(ctx context.Context, agentPool string, options AgentPoolUpdateOptions) (*AgentPool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, agentPool, options)
	ret0, _ := ret[0].(*AgentPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAgentPoolsMockRecorder) Update(ctx, agentPool, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAgentPools)(nil).Update), ctx, agentPool, options)
}