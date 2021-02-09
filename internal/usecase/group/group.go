package group

import (
	"context"

	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	GroupGateway interface { //nolint:golint
		FindGroupByID(ctx context.Context, id int64) (model.Group, error)
		FindAllGroups(ctx context.Context) ([]model.Group, error)
		CreateGroup(ctx context.Context, group model.Group) (int64, error)
		DeleteGroupByID(ctx context.Context, id int64) (int64, error)
		UpdateGroup(ctx context.Context, group model.Group) (int64, error)
	}

	UseCase struct {
		groupGateway GroupGateway
	}
)

func NewUseCase(groupGateway GroupGateway) *UseCase {
	return &UseCase{groupGateway: groupGateway}
}

func (uc *UseCase) CreateGroup(ctx context.Context, group model.Group) (int64, error) {
	return uc.groupGateway.CreateGroup(ctx, group)
}

func (uc *UseCase) UpdateGroup(ctx context.Context, group model.Group) (int64, error) {
	return uc.groupGateway.UpdateGroup(ctx, group)
}

func (uc *UseCase) DeleteGroupByID(ctx context.Context, id int64) (int64, error) {
	return uc.groupGateway.DeleteGroupByID(ctx, id)
}

func (uc *UseCase) FindGroupByID(ctx context.Context, id int64) (model.Group, error) {
	return uc.groupGateway.FindGroupByID(ctx, id)
}

func (uc *UseCase) FindAllGroups(ctx context.Context) ([]model.Group, error) {
	return uc.groupGateway.FindAllGroups(ctx)
}
