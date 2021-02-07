// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	model "github.com/sterligov/banner-rotator/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// GatewayEvent is an autogenerated mock type for the GatewayEvent type
type GatewayEvent struct {
	mock.Mock
}

// Publish provides a mock function with given fields: e
func (_m *GatewayEvent) Publish(e model.Event) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Event) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
