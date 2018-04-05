// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mulesoft-labs/aws-keycloak/provider (interfaces: AwsProviderIf)

// Package mock_provider is a generated GoMock package.
package mock_provider

import (
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAwsProviderIf is a mock of AwsProviderIf interface
type MockAwsProviderIf struct {
	ctrl     *gomock.Controller
	recorder *MockAwsProviderIfMockRecorder
}

// MockAwsProviderIfMockRecorder is the mock recorder for MockAwsProviderIf
type MockAwsProviderIfMockRecorder struct {
	mock *MockAwsProviderIf
}

// NewMockAwsProviderIf creates a new mock instance
func NewMockAwsProviderIf(ctrl *gomock.Controller) *MockAwsProviderIf {
	mock := &MockAwsProviderIf{ctrl: ctrl}
	mock.recorder = &MockAwsProviderIfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAwsProviderIf) EXPECT() *MockAwsProviderIfMockRecorder {
	return m.recorder
}

// AssumeRoleWithSAML mocks base method
func (m *MockAwsProviderIf) AssumeRoleWithSAML(arg0, arg1, arg2 string) (sts.Credentials, error) {
	ret := m.ctrl.Call(m, "AssumeRoleWithSAML", arg0, arg1, arg2)
	ret0, _ := ret[0].(sts.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssumeRoleWithSAML indicates an expected call of AssumeRoleWithSAML
func (mr *MockAwsProviderIfMockRecorder) AssumeRoleWithSAML(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssumeRoleWithSAML", reflect.TypeOf((*MockAwsProviderIf)(nil).AssumeRoleWithSAML), arg0, arg1, arg2)
}

// CheckAlreadyAuthd mocks base method
func (m *MockAwsProviderIf) CheckAlreadyAuthd(arg0 string) (sts.Credentials, error) {
	ret := m.ctrl.Call(m, "CheckAlreadyAuthd", arg0)
	ret0, _ := ret[0].(sts.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckAlreadyAuthd indicates an expected call of CheckAlreadyAuthd
func (mr *MockAwsProviderIfMockRecorder) CheckAlreadyAuthd(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAlreadyAuthd", reflect.TypeOf((*MockAwsProviderIf)(nil).CheckAlreadyAuthd), arg0)
}

// StoreAwsCreds mocks base method
func (m *MockAwsProviderIf) StoreAwsCreds(arg0 sts.Credentials, arg1 string) {
	m.ctrl.Call(m, "StoreAwsCreds", arg0, arg1)
}

// StoreAwsCreds indicates an expected call of StoreAwsCreds
func (mr *MockAwsProviderIfMockRecorder) StoreAwsCreds(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreAwsCreds", reflect.TypeOf((*MockAwsProviderIf)(nil).StoreAwsCreds), arg0, arg1)
}
