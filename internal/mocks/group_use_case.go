// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/sterligov/banner-rotator/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// GroupUseCase is an autogenerated mock type for the GroupUseCase type
type GroupUseCase struct {
	mock.Mock
}

// CreateGroup provides a mock function with given fields: ctx, b
func (_m *GroupUseCase) CreateGroup(ctx context.Context, b model.Group) (int64, error) {
	ret := _m.Called(ctx, b)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.Group) int64); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Group) error); ok {
		r1 = rf(ctx, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteGroupByID provides a mock function with given fields: ctx, id
func (_m *GroupUseCase) DeleteGroupByID(ctx context.Context, id int64) (int64, error) {
	ret := _m.Called(ctx, id)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64) int64); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllGroups provides a mock function with given fields: ctx
func (_m *GroupUseCase) FindAllGroups(ctx context.Context) ([]model.Group, error) {
	ret := _m.Called(ctx)

	var r0 []model.Group
	if rf, ok := ret.Get(0).(func(context.Context) []model.Group); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindGroupByID provides a mock function with given fields: ctx, id
func (_m *GroupUseCase) FindGroupByID(ctx context.Context, id int64) (model.Group, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Group
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Group); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Group)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateGroup provides a mock function with given fields: ctx, b
func (_m *GroupUseCase) UpdateGroup(ctx context.Context, b model.Group) (int64, error) {
	ret := _m.Called(ctx, b)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.Group) int64); ok {
		r0 = rf(ctx, b)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Group) error); ok {
		r1 = rf(ctx, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
