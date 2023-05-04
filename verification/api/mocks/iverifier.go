// Code generated by MockGen. DO NOT EDIT.
// Source: ../verifier/iverifier.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/veraison/services/proto"
)

// MockIVerifier is a mock of IVerifier interface.
type MockIVerifier struct {
	ctrl     *gomock.Controller
	recorder *MockIVerifierMockRecorder
}

// MockIVerifierMockRecorder is the mock recorder for MockIVerifier.
type MockIVerifierMockRecorder struct {
	mock *MockIVerifier
}

// NewMockIVerifier creates a new mock instance.
func NewMockIVerifier(ctrl *gomock.Controller) *MockIVerifier {
	mock := &MockIVerifier{ctrl: ctrl}
	mock.recorder = &MockIVerifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIVerifier) EXPECT() *MockIVerifierMockRecorder {
	return m.recorder
}

// GetPublicKey mocks base method.
func (m *MockIVerifier) GetPublicKey() (*proto.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey")
	ret0, _ := ret[0].(*proto.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockIVerifierMockRecorder) GetPublicKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockIVerifier)(nil).GetPublicKey))
}

// GetVTSState mocks base method.
func (m *MockIVerifier) GetVTSState() (*proto.ServiceState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVTSState")
	ret0, _ := ret[0].(*proto.ServiceState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVTSState indicates an expected call of GetVTSState.
func (mr *MockIVerifierMockRecorder) GetVTSState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVTSState", reflect.TypeOf((*MockIVerifier)(nil).GetVTSState))
}

// IsSupportedMediaType mocks base method.
func (m *MockIVerifier) IsSupportedMediaType(mt string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSupportedMediaType", mt)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSupportedMediaType indicates an expected call of IsSupportedMediaType.
func (mr *MockIVerifierMockRecorder) IsSupportedMediaType(mt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSupportedMediaType", reflect.TypeOf((*MockIVerifier)(nil).IsSupportedMediaType), mt)
}

// ProcessEvidence mocks base method.
func (m *MockIVerifier) ProcessEvidence(tenantID string, nonce, data []byte, mt string, teeReport bool) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessEvidence", tenantID, nonce, data, mt, teeReport)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessEvidence indicates an expected call of ProcessEvidence.
func (mr *MockIVerifierMockRecorder) ProcessEvidence(tenantID, nonce, data, mt, teeReport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessEvidence", reflect.TypeOf((*MockIVerifier)(nil).ProcessEvidence), tenantID, nonce, data, mt, teeReport)
}

// SupportedMediaTypes mocks base method.
func (m *MockIVerifier) SupportedMediaTypes() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportedMediaTypes")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SupportedMediaTypes indicates an expected call of SupportedMediaTypes.
func (mr *MockIVerifierMockRecorder) SupportedMediaTypes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedMediaTypes", reflect.TypeOf((*MockIVerifier)(nil).SupportedMediaTypes))
}
