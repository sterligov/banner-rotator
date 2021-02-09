package group

import (
	"context"
	"fmt"
	"testing"

	"github.com/sterligov/banner-rotator/internal/mocks"
	"github.com/sterligov/banner-rotator/internal/model"
	"github.com/stretchr/testify/require"
)

func TestFindAllGroups(t *testing.T) {
	tests := []struct {
		groups []model.Group
		err    error
		name   string
	}{
		{[]model.Group{{ID: 1, Description: "descr"}, {ID: 2, Description: "descr2"}}, nil, "ok"},
		{[]model.Group{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupGw := &mocks.GroupGateway{}

			ctx := context.Background()
			groupGw.
				On("FindAll", ctx).
				Return(tst.groups, tst.err).
				Once()
			defer groupGw.AssertExpectations(t)

			uc := NewUseCase(groupGw)
			groups, err := uc.FindAllGroups(ctx)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.groups, groups)
		})
	}
}

func TestFindGroupByID(t *testing.T) {
	tests := []struct {
		group model.Group
		err   error
		name  string
	}{
		{model.Group{ID: 1, Description: "descr"}, nil, "ok"},
		{model.Group{}, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupGw := &mocks.GroupGateway{}

			var groupID int64 = 1
			ctx := context.Background()
			groupGw.
				On("FindByID", ctx, groupID).
				Return(tst.group, tst.err).
				Once()
			defer groupGw.AssertExpectations(t)

			uc := NewUseCase(groupGw)
			group, err := uc.FindGroupByID(ctx, groupID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.group, group)
		})
	}
}

func TestCreateGroup(t *testing.T) {
	tests := []struct {
		insertedID int64
		err        error
		name       string
	}{
		{1, nil, "ok"},
		{0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupGw := &mocks.GroupGateway{}

			group := model.Group{
				Description: "description",
			}

			ctx := context.Background()
			groupGw.
				On("Create", ctx, group).
				Return(tst.insertedID, tst.err).
				Once()
			defer groupGw.AssertExpectations(t)

			uc := NewUseCase(groupGw)
			insertedID, err := uc.CreateGroup(ctx, group)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.insertedID, insertedID)
		})
	}
}

func TestUpdateGroup(t *testing.T) {
	tests := []struct {
		affected int64
		err      error
		name     string
	}{
		{1, nil, "ok"},
		{0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupGw := &mocks.GroupGateway{}

			group := model.Group{
				Description: "description",
			}

			ctx := context.Background()
			groupGw.
				On("Update", ctx, group).
				Return(tst.affected, tst.err).
				Once()
			defer groupGw.AssertExpectations(t)

			uc := NewUseCase(groupGw)
			insertedID, err := uc.UpdateGroup(ctx, group)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, insertedID)
		})
	}
}

func TestDeleteGroupByID(t *testing.T) {
	tests := []struct {
		affected int64
		err      error
		name     string
	}{
		{1, nil, "ok"},
		{0, fmt.Errorf("error"), "error"},
	}

	for _, tst := range tests {
		tst := tst
		t.Run(tst.name, func(t *testing.T) {
			groupGw := &mocks.GroupGateway{}

			var groupID int64 = 1

			ctx := context.Background()
			groupGw.
				On("DeleteByID", ctx, groupID).
				Return(tst.affected, tst.err).
				Once()
			defer groupGw.AssertExpectations(t)

			uc := NewUseCase(groupGw)
			affected, err := uc.DeleteGroupByID(ctx, groupID)
			require.Equal(t, tst.err, err)
			require.Equal(t, tst.affected, affected)
		})
	}
}
