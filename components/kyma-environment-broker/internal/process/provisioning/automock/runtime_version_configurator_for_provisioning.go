// Code generated by mockery v2.3.0. DO NOT EDIT.

package automock

import (
	internal "github.com/kyma-project/control-plane/components/kyma-environment-broker/internal"
	mock "github.com/stretchr/testify/mock"
)

// RuntimeVersionConfiguratorForProvisioning is an autogenerated mock type for the RuntimeVersionConfiguratorForProvisioning type
type RuntimeVersionConfiguratorForProvisioning struct {
	mock.Mock
}

// ForProvisioning provides a mock function with given fields: op, pp
func (_m *RuntimeVersionConfiguratorForProvisioning) ForProvisioning(op internal.ProvisioningOperation, pp internal.ProvisioningParameters) (*internal.RuntimeVersionData, error) {
	ret := _m.Called(op, pp)

	var r0 *internal.RuntimeVersionData
	if rf, ok := ret.Get(0).(func(internal.ProvisioningOperation, internal.ProvisioningParameters) *internal.RuntimeVersionData); ok {
		r0 = rf(op, pp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*internal.RuntimeVersionData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(internal.ProvisioningOperation, internal.ProvisioningParameters) error); ok {
		r1 = rf(op, pp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
