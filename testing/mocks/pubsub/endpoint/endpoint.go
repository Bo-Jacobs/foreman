// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-foreman/foreman/pubsub/endpoint (interfaces: Endpoint)

// Package endpoint is a generated GoMock package.
package endpoint

import (
	context "context"
	reflect "reflect"

	endpoint "github.com/go-foreman/foreman/pubsub/endpoint"
	message "github.com/go-foreman/foreman/pubsub/message"
	gomock "github.com/golang/mock/gomock"
)

// MockEndpoint is a mock of Endpoint interface.
type MockEndpoint struct {
	ctrl     *gomock.Controller
	recorder *MockEndpointMockRecorder
}

// MockEndpointMockRecorder is the mock recorder for MockEndpoint.
type MockEndpointMockRecorder struct {
	mock *MockEndpoint
}

// NewMockEndpoint creates a new mock instance.
func NewMockEndpoint(ctrl *gomock.Controller) *MockEndpoint {
	mock := &MockEndpoint{ctrl: ctrl}
	mock.recorder = &MockEndpointMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEndpoint) EXPECT() *MockEndpointMockRecorder {
	return m.recorder
}

// Name mocks base method.
func (m *MockEndpoint) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockEndpointMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockEndpoint)(nil).Name))
}

// Send mocks base method.
func (m *MockEndpoint) Send(arg0 context.Context, arg1 *message.OutcomingMessage, arg2 ...endpoint.DeliveryOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Send", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockEndpointMockRecorder) Send(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockEndpoint)(nil).Send), varargs...)
}