package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/weblogs/app/blog/domain/repository"
	repoimpl "github.com/weblogs/app/blog/infra/repository"
	"github.com/weblogs/app/blog/model/entity"
	"github.com/weblogs/pkg/idgenerator"
)

type PostUpdateService struct {
	repo repository.BlogPostRepo
}

var postUpdateSrv *PostUpdateService

func GetBlogUpdateServiceInstance() *PostUpdateService {
	if postUpdateSrv == nil {
		postUpdateSrv = NewPostUpdateService(repoimpl.NewBlogPostRepo())
	}
	return postUpdateSrv
}

func NewPostUpdateService(repo repository.BlogPostRepo) *PostUpdateService {
	return &PostUpdateService{
		repo: repo,
	}
}

func (s *PostUpdateService) CreatePost(ctx context.Context, e *entity.BlogPost) error {

	e.PostID = idgenerator.GetIDGeneratorInstance().ID()
	e.Model = entity.NewModel()
	e.Status = entity.PostCreated
	err := s.repo.Create(ctx, e)
	if err != nil {
		return errors.Wrapf(err, "PostUpdateService::CreatePost failed")
	}
	return nil
}

func (s *PostUpdateService) UpdatePost(ctx context.Context, e *entity.BlogPost) error {
	s.repo.Update(ctx)
}
