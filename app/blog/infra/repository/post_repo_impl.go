package repository

import (
	"context"
	"errors"

	"github.com/weblogs/app/blog/domain/repository"
	"github.com/weblogs/app/blog/infra/repository/convert"
	"github.com/weblogs/app/blog/model/entity"
	"github.com/weblogs/pkg/conf"
	"github.com/weblogs/pkg/mongo"
)

type BlogPostRepoImpl struct{}

func NewBlogPostRepo() repository.BlogPostRepo {
	return BlogPostRepoImpl{}
}

func (b BlogPostRepoImpl) Create(ctx context.Context, e *entity.BlogPost) error {
	if e == nil {
		return errors.New("insert data should not be empty")
	}

	po, err := convert.PostEntity2POConverter.ToPO(ctx, e)
	if err != nil {
		return err
	}

	coll := mongo.NewMongoDB(conf.DBName()).Collection(e.CollectionName())
	_, err = coll.InsertOne(ctx, po)
	if err != nil {
		return err
	}

	return nil
}

func (b BlogPostRepoImpl) Update(ctx context.Context, origin, target *entity.BlogPost) error {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostRepoImpl) Get(ctx context.Context, postID int64) (*entity.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostRepoImpl) MGet(ctx context.Context, postIDs []int64) ([]*entity.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostRepoImpl) Find(ctx context.Context, filter map[string]any) ([]*entity.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostRepoImpl) FindOne(ctx context.Context, filter map[string]any) (*entity.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}
