package repository

import (
	"context"

	"github.com/weblogs/app/user/model/entity"
	"github.com/weblogs/app/user/model/po/field"
)

type UserRepo interface {
	Create(ctx context.Context, e *entity.User) error
	Query(ctx context.Context, filters map[field.Field]any, projections ...field.Field) ([]*entity.User, error)
	Get(ctx context.Context, userID int64) (*entity.User, error)
	MGet(ctx context.Context, userIDs []int64) ([]*entity.User, error)
	Update(ctx context.Context, e *entity.User) error
	Delete(ctx context.Context, e *entity.User) error
}

func NewUserRepo() UserRepo {
	return nil
}

var (
	userRepo = NewUserRepo()
)

func GetUserRepoInstance() UserRepo {
	return userRepo
}
