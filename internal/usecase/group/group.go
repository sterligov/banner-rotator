package group

import (
	"context"

	"github.com/sterligov/banner-rotator/internal/model"
)

type (
	Gateway interface {
		FindByID(ctx context.Context, id int64) (model.Group, error)
		FindAll(ctx context.Context) ([]model.Group, error)
		Create(ctx context.Context, group model.Group) (int64, error)
		DeleteByID(ctx context.Context, id int64) (int64, error)
		Update(ctx context.Context, group model.Group) (int64, error)
	}

	UseCase struct {
		groupGateway Gateway
	}
)

func NewUseCase(groupGateway Gateway) *UseCase {
	return &UseCase{groupGateway: groupGateway}
}

func (uc *UseCase) CreateGroup(ctx context.Context, group model.Group) (int64, error) {
	return uc.groupGateway.Create(ctx, group)
}

func (uc *UseCase) UpdateGroup(ctx context.Context, group model.Group) (int64, error) {
	return uc.groupGateway.Update(ctx, group)
}

func (uc *UseCase) DeleteGroupByID(ctx context.Context, id int64) (int64, error) {
	return uc.groupGateway.DeleteByID(ctx, id)
}

func (uc *UseCase) FindGroupByID(ctx context.Context, id int64) (model.Group, error) {
	return uc.groupGateway.FindByID(ctx, id)
}

func (uc *UseCase) FindAllGroups(ctx context.Context) ([]model.Group, error) {
	return uc.groupGateway.FindAll(ctx)
}
