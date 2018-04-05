// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mulesoft-labs/aws-keycloak/provider (interfaces: KeycloakProviderIf)

// Package mock_provider is a generated GoMock package.
package mock_provider

import (
	gomock "github.com/golang/mock/gomock"
	provider "github.com/mulesoft-labs/aws-keycloak/provider"
	reflect "reflect"
)

// MockKeycloakProviderIf is a mock of KeycloakProviderIf interface
type MockKeycloakProviderIf struct {
	ctrl     *gomock.Controller
	recorder *MockKeycloakProviderIfMockRecorder
}

// MockKeycloakProviderIfMockRecorder is the mock recorder for MockKeycloakProviderIf
type MockKeycloakProviderIfMockRecorder struct {
	mock *MockKeycloakProviderIf
}

// NewMockKeycloakProviderIf creates a new mock instance
func NewMockKeycloakProviderIf(ctrl *gomock.Controller) *MockKeycloakProviderIf {
	mock := &MockKeycloakProviderIf{ctrl: ctrl}
	mock.recorder = &MockKeycloakProviderIfMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKeycloakProviderIf) EXPECT() *MockKeycloakProviderIfMockRecorder {
	return m.recorder
}

// BasicAuth mocks base method
func (m *MockKeycloakProviderIf) BasicAuth() error {
	ret := m.ctrl.Call(m, "BasicAuth")
	ret0, _ := ret[0].(error)
	return ret0
}

// BasicAuth indicates an expected call of BasicAuth
func (mr *MockKeycloakProviderIfMockRecorder) BasicAuth() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BasicAuth", reflect.TypeOf((*MockKeycloakProviderIf)(nil).BasicAuth))
}

// GetSamlAssertion mocks base method
func (m *MockKeycloakProviderIf) GetSamlAssertion() (provider.SAMLAssertion, error) {
	ret := m.ctrl.Call(m, "GetSamlAssertion")
	ret0, _ := ret[0].(provider.SAMLAssertion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSamlAssertion indicates an expected call of GetSamlAssertion
func (mr *MockKeycloakProviderIfMockRecorder) GetSamlAssertion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSamlAssertion", reflect.TypeOf((*MockKeycloakProviderIf)(nil).GetSamlAssertion))
}

// RetrieveKeycloakCreds mocks base method
func (m *MockKeycloakProviderIf) RetrieveKeycloakCreds() bool {
	ret := m.ctrl.Call(m, "RetrieveKeycloakCreds")
	ret0, _ := ret[0].(bool)
	return ret0
}

// RetrieveKeycloakCreds indicates an expected call of RetrieveKeycloakCreds
func (mr *MockKeycloakProviderIfMockRecorder) RetrieveKeycloakCreds() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RetrieveKeycloakCreds", reflect.TypeOf((*MockKeycloakProviderIf)(nil).RetrieveKeycloakCreds))
}

// StoreKeycloakCreds mocks base method
func (m *MockKeycloakProviderIf) StoreKeycloakCreds() {
	m.ctrl.Call(m, "StoreKeycloakCreds")
}

// StoreKeycloakCreds indicates an expected call of StoreKeycloakCreds
func (mr *MockKeycloakProviderIfMockRecorder) StoreKeycloakCreds() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreKeycloakCreds", reflect.TypeOf((*MockKeycloakProviderIf)(nil).StoreKeycloakCreds))
}
