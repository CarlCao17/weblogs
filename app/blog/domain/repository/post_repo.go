package repository

import (
	"context"

	"github.com/weblogs/app/blog/model/entity"
)

type BlogPostRepo interface {
	Create(ctx context.Context, e *entity.BlogPost) error
	Update(ctx context.Context, origin, target *entity.BlogPost) error
	Get(ctx context.Context, postID int64) (*entity.BlogPost, error)
	MGet(ctx context.Context, postIDs []int64) ([]*entity.BlogPost, error)
	Find(ctx context.Context, filter map[string]any) ([]*entity.BlogPost, error)
	FindOne(ctx context.Context, filter map[string]any) (*entity.BlogPost, error)
}
