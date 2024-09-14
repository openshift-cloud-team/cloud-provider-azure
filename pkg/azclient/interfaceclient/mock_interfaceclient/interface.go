// /*
// Copyright The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// Code generated by MockGen. DO NOT EDIT.
// Source: interfaceclient/interface.go
//
// Generated by this command:
//
//	mockgen -package mock_interfaceclient -source interfaceclient/interface.go
//

// Package mock_interfaceclient is a generated GoMock package.
package mock_interfaceclient

import (
	context "context"
	reflect "reflect"

	armnetwork "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v6"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockInterface) CreateOrUpdate(ctx context.Context, resourceGroupName, resourceName string, resourceParam armnetwork.Interface) (*armnetwork.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", ctx, resourceGroupName, resourceName, resourceParam)
	ret0, _ := ret[0].(*armnetwork.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockInterfaceMockRecorder) CreateOrUpdate(ctx, resourceGroupName, resourceName, resourceParam any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdate), ctx, resourceGroupName, resourceName, resourceParam)
}

// Delete mocks base method.
func (m *MockInterface) Delete(ctx context.Context, resourceGroupName, resourceName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, resourceGroupName, resourceName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockInterfaceMockRecorder) Delete(ctx, resourceGroupName, resourceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInterface)(nil).Delete), ctx, resourceGroupName, resourceName)
}

// Get mocks base method.
func (m *MockInterface) Get(ctx context.Context, resourceGroupName, resourceName string, expand *string) (*armnetwork.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, resourceGroupName, resourceName, expand)
	ret0, _ := ret[0].(*armnetwork.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(ctx, resourceGroupName, resourceName, expand any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), ctx, resourceGroupName, resourceName, expand)
}

// GetVirtualMachineScaleSetNetworkInterface mocks base method.
func (m *MockInterface) GetVirtualMachineScaleSetNetworkInterface(ctx context.Context, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName string) (*armnetwork.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMachineScaleSetNetworkInterface", ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName)
	ret0, _ := ret[0].(*armnetwork.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMachineScaleSetNetworkInterface indicates an expected call of GetVirtualMachineScaleSetNetworkInterface.
func (mr *MockInterfaceMockRecorder) GetVirtualMachineScaleSetNetworkInterface(ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMachineScaleSetNetworkInterface", reflect.TypeOf((*MockInterface)(nil).GetVirtualMachineScaleSetNetworkInterface), ctx, resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName)
}

// List mocks base method.
func (m *MockInterface) List(ctx context.Context, resourceGroupName string) ([]*armnetwork.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, resourceGroupName)
	ret0, _ := ret[0].([]*armnetwork.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockInterfaceMockRecorder) List(ctx, resourceGroupName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInterface)(nil).List), ctx, resourceGroupName)
}

// ListVirtualMachineScaleSetNetworkInterfaces mocks base method.
func (m *MockInterface) ListVirtualMachineScaleSetNetworkInterfaces(ctx context.Context, resourceGroupName, virtualMachineScaleSetName string) ([]*armnetwork.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVirtualMachineScaleSetNetworkInterfaces", ctx, resourceGroupName, virtualMachineScaleSetName)
	ret0, _ := ret[0].([]*armnetwork.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualMachineScaleSetNetworkInterfaces indicates an expected call of ListVirtualMachineScaleSetNetworkInterfaces.
func (mr *MockInterfaceMockRecorder) ListVirtualMachineScaleSetNetworkInterfaces(ctx, resourceGroupName, virtualMachineScaleSetName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualMachineScaleSetNetworkInterfaces", reflect.TypeOf((*MockInterface)(nil).ListVirtualMachineScaleSetNetworkInterfaces), ctx, resourceGroupName, virtualMachineScaleSetName)
}
