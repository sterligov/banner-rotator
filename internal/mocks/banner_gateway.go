// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/sterligov/banner-rotator/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// BannerGateway is an autogenerated mock type for the BannerGateway type
type BannerGateway struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *BannerGateway) Create(ctx context.Context, _a1 model.Banner) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.Banner) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Banner) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateBannerSlotRelation provides a mock function with given fields: ctx, bannerID, slotID
func (_m *BannerGateway) CreateBannerSlotRelation(ctx context.Context, bannerID int64, slotID int64) (int64, error) {
	ret := _m.Called(ctx, bannerID, slotID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) int64); ok {
		r0 = rf(ctx, bannerID, slotID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, bannerID, slotID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBannerSlotRelation provides a mock function with given fields: ctx, bannerID, slotID
func (_m *BannerGateway) DeleteBannerSlotRelation(ctx context.Context, bannerID int64, slotID int64) (int64, error) {
	ret := _m.Called(ctx, bannerID, slotID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) int64); ok {
		r0 = rf(ctx, bannerID, slotID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int64) error); ok {
		r1 = rf(ctx, bannerID, slotID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *BannerGateway) DeleteByID(ctx context.Context, id int64) (int64, error) {
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

// FindAll provides a mock function with given fields: ctx
func (_m *BannerGateway) FindAll(ctx context.Context) ([]model.Banner, error) {
	ret := _m.Called(ctx)

	var r0 []model.Banner
	if rf, ok := ret.Get(0).(func(context.Context) []model.Banner); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Banner)
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

// FindByID provides a mock function with given fields: ctx, id
func (_m *BannerGateway) FindByID(ctx context.Context, id int64) (model.Banner, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Banner
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Banner); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Banner)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *BannerGateway) Update(ctx context.Context, _a1 model.Banner) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.Banner) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Banner) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
