// Code generated by MockGen. DO NOT EDIT.
// Source: metrics.go

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	prometheus "github.com/prometheus/client_golang/prometheus"
	dynamic "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/http/config/dynamic"
	metrics "gitlab.com/ConsenSys/client/fr/core-stack/orchestrate.git/v2/pkg/metrics"
	reflect "reflect"
)

// MockPrometheus is a mock of Prometheus interface
type MockPrometheus struct {
	ctrl     *gomock.Controller
	recorder *MockPrometheusMockRecorder
}

// MockPrometheusMockRecorder is the mock recorder for MockPrometheus
type MockPrometheusMockRecorder struct {
	mock *MockPrometheus
}

// NewMockPrometheus creates a new mock instance
func NewMockPrometheus(ctrl *gomock.Controller) *MockPrometheus {
	mock := &MockPrometheus{ctrl: ctrl}
	mock.recorder = &MockPrometheusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPrometheus) EXPECT() *MockPrometheusMockRecorder {
	return m.recorder
}

// Describe mocks base method
func (m *MockPrometheus) Describe(arg0 chan<- *prometheus.Desc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Describe", arg0)
}

// Describe indicates an expected call of Describe
func (mr *MockPrometheusMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockPrometheus)(nil).Describe), arg0)
}

// Collect mocks base method
func (m *MockPrometheus) Collect(arg0 chan<- prometheus.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Collect", arg0)
}

// Collect indicates an expected call of Collect
func (mr *MockPrometheusMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockPrometheus)(nil).Collect), arg0)
}

// MockDynamicPrometheus is a mock of DynamicPrometheus interface
type MockDynamicPrometheus struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicPrometheusMockRecorder
}

// MockDynamicPrometheusMockRecorder is the mock recorder for MockDynamicPrometheus
type MockDynamicPrometheusMockRecorder struct {
	mock *MockDynamicPrometheus
}

// NewMockDynamicPrometheus creates a new mock instance
func NewMockDynamicPrometheus(ctrl *gomock.Controller) *MockDynamicPrometheus {
	mock := &MockDynamicPrometheus{ctrl: ctrl}
	mock.recorder = &MockDynamicPrometheusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamicPrometheus) EXPECT() *MockDynamicPrometheusMockRecorder {
	return m.recorder
}

// Switch mocks base method
func (m *MockDynamicPrometheus) Switch(arg0 *dynamic.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Switch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Switch indicates an expected call of Switch
func (mr *MockDynamicPrometheusMockRecorder) Switch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Switch", reflect.TypeOf((*MockDynamicPrometheus)(nil).Switch), arg0)
}

// Describe mocks base method
func (m *MockDynamicPrometheus) Describe(arg0 chan<- *prometheus.Desc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Describe", arg0)
}

// Describe indicates an expected call of Describe
func (mr *MockDynamicPrometheusMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockDynamicPrometheus)(nil).Describe), arg0)
}

// Collect mocks base method
func (m *MockDynamicPrometheus) Collect(arg0 chan<- prometheus.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Collect", arg0)
}

// Collect indicates an expected call of Collect
func (mr *MockDynamicPrometheusMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockDynamicPrometheus)(nil).Collect), arg0)
}

// MockRegistry is a mock of Registry interface
type MockRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockRegistryMockRecorder
}

// MockRegistryMockRecorder is the mock recorder for MockRegistry
type MockRegistryMockRecorder struct {
	mock *MockRegistry
}

// NewMockRegistry creates a new mock instance
func NewMockRegistry(ctrl *gomock.Controller) *MockRegistry {
	mock := &MockRegistry{ctrl: ctrl}
	mock.recorder = &MockRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegistry) EXPECT() *MockRegistryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m_2 *MockRegistry) Add(m metrics.Prometheus) {
	m_2.ctrl.T.Helper()
	m_2.ctrl.Call(m_2, "Add", m)
}

// Add indicates an expected call of Add
func (mr *MockRegistryMockRecorder) Add(m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRegistry)(nil).Add), m)
}

// SwitchDynConfig mocks base method
func (m *MockRegistry) SwitchDynConfig(dynCfg *dynamic.Configuration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SwitchDynConfig", dynCfg)
	ret0, _ := ret[0].(error)
	return ret0
}

// SwitchDynConfig indicates an expected call of SwitchDynConfig
func (mr *MockRegistryMockRecorder) SwitchDynConfig(dynCfg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SwitchDynConfig", reflect.TypeOf((*MockRegistry)(nil).SwitchDynConfig), dynCfg)
}

// Describe mocks base method
func (m *MockRegistry) Describe(arg0 chan<- *prometheus.Desc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Describe", arg0)
}

// Describe indicates an expected call of Describe
func (mr *MockRegistryMockRecorder) Describe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Describe", reflect.TypeOf((*MockRegistry)(nil).Describe), arg0)
}

// Collect mocks base method
func (m *MockRegistry) Collect(arg0 chan<- prometheus.Metric) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Collect", arg0)
}

// Collect indicates an expected call of Collect
func (mr *MockRegistryMockRecorder) Collect(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collect", reflect.TypeOf((*MockRegistry)(nil).Collect), arg0)
}