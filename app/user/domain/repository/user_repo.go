package repository

import (
	"context"

	"github.com/weblogs/app/user/model/entity"
)

type UserRepo interface {
	Create(ctx context.Context, e *entity.User) error
	Query(ctx context.Context, filterParam map[string]any) ([]*entity.User, error)
	Get(ctx context.Context, userID int64) (*entity.User, error)
	MGet(ctx context.Context, userIDs []int64) ([]*entity.User, error)
	Update(ctx context.Context, e *entity.User) error
	Delete(ctx context.Context, e *entity.User) error
}
