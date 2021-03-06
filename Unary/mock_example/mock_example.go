// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/freddy50806/gRPCTestWithGomock/Unary/proto (interfaces: ExampleServiceClient)

// Package mock_proto is a generated GoMock package.
package mock_proto

import (
	context "context"
	proto "github.com/freddy50806/gRPCTestWithGomock/Unary/proto"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockExampleServiceClient is a mock of ExampleServiceClient interface
type MockExampleServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockExampleServiceClientMockRecorder
}

// MockExampleServiceClientMockRecorder is the mock recorder for MockExampleServiceClient
type MockExampleServiceClientMockRecorder struct {
	mock *MockExampleServiceClient
}

// NewMockExampleServiceClient creates a new mock instance
func NewMockExampleServiceClient(ctrl *gomock.Controller) *MockExampleServiceClient {
	mock := &MockExampleServiceClient{ctrl: ctrl}
	mock.recorder = &MockExampleServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExampleServiceClient) EXPECT() *MockExampleServiceClientMockRecorder {
	return m.recorder
}

// IsEven mocks base method
func (m *MockExampleServiceClient) IsEven(arg0 context.Context, arg1 *proto.Request, arg2 ...grpc.CallOption) (*proto.Reply, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsEven", varargs...)
	ret0, _ := ret[0].(*proto.Reply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsEven indicates an expected call of IsEven
func (mr *MockExampleServiceClientMockRecorder) IsEven(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsEven", reflect.TypeOf((*MockExampleServiceClient)(nil).IsEven), varargs...)
}
