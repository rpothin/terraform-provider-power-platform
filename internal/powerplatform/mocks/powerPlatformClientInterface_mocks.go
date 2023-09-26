// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api/ppapi (interfaces: PowerPlatformClientInterface)

// Package powerplatform_mocks is a generated GoMock package.
package powerplatform_mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	powerplatform_common "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/api"
	powerplatform_models "github.com/microsoft/terraform-provider-power-platform/internal/powerplatform/models"
)

// MockPowerPlatformClientInterface is a mock of PowerPlatformClientInterface interface.
type MockPowerPlatformClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPowerPlatformClientInterfaceMockRecorder
}

// MockPowerPlatformClientInterfaceMockRecorder is the mock recorder for MockPowerPlatformClientInterface.
type MockPowerPlatformClientInterfaceMockRecorder struct {
	mock *MockPowerPlatformClientInterface
}

// NewMockPowerPlatformClientInterface creates a new mock instance.
func NewMockPowerPlatformClientInterface(ctrl *gomock.Controller) *MockPowerPlatformClientInterface {
	mock := &MockPowerPlatformClientInterface{ctrl: ctrl}
	mock.recorder = &MockPowerPlatformClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPowerPlatformClientInterface) EXPECT() *MockPowerPlatformClientInterfaceMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockPowerPlatformClientInterface) Execute(arg0 context.Context, arg1, arg2 string, arg3 interface{}, arg4 []int, arg5 interface{}) (*powerplatform_common.ApiHttpResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(*powerplatform_common.ApiHttpResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockPowerPlatformClientInterfaceMockRecorder) Execute(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockPowerPlatformClientInterface)(nil).Execute), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetBase mocks base method.
func (m *MockPowerPlatformClientInterface) GetBase() powerplatform_common.ApiClientInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBase")
	ret0, _ := ret[0].(powerplatform_common.ApiClientInterface)
	return ret0
}

// GetBase indicates an expected call of GetBase.
func (mr *MockPowerPlatformClientInterfaceMockRecorder) GetBase() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBase", reflect.TypeOf((*MockPowerPlatformClientInterface)(nil).GetBase))
}

// GetBillingPolicies mocks base method.
func (m *MockPowerPlatformClientInterface) GetBillingPolicies(arg0 context.Context) ([]powerplatform_models.BillingPolicyDto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBillingPolicies", arg0)
	ret0, _ := ret[0].([]powerplatform_models.BillingPolicyDto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBillingPolicies indicates an expected call of GetBillingPolicies.
func (mr *MockPowerPlatformClientInterfaceMockRecorder) GetBillingPolicies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBillingPolicies", reflect.TypeOf((*MockPowerPlatformClientInterface)(nil).GetBillingPolicies), arg0)
}