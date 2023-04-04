package service

import (
	"context"

	"github.com/weblogs/app/user/domain/repository"
	"github.com/weblogs/app/user/model/entity"
)

type UserQuery struct {
	repo repository.UserRepo
}

func GetUserQueryServiceInstance() *UserQuery {
	if userQueryService == nil {
		userQueryService = NewUserQueryService()
	}
	return userQueryService
}

var userQueryService *UserQuery

func NewUserQueryService(repo repository.UserRepo) *UserQuery {
	return &UserQuery{
		repo: repo,
	}
}

func (s *UserQuery) Get(ctx context.Context, userID int64) (*entity.User, error) {
	panic("")
}

func (s *UserQuery) MGet(ctx context.Context, userIDs []int64) ([]*entity.User, error) {
	panic("")
}

func (s *UserQuery) Check(ctx context.Context, userName, passwd string) error {
	panic("")
}

func (s *UserQuery) Exists(ctx context.Context, userName string) bool {
	panic("")
}

func (s *UserQuery) Search(ctx context.Context, filterParams map[string]any, fields []string) ([]*entity.User, error) {
	panic("")
}

func (s *UserQuery) SearchByUserName(ctx context.Context, name string) (*entity.User, error) {
	panic("")
}

func (s *UserQuery) SearchByUserRole(ctx context.Context, role string) ([]*entity.User, error) {
	panic("")
}

func (s *UserQuery) SearchByEmail(ctx context.Context, email string) (*entity.User, error) {
	panic("")
}

func (s *UserQuery) SearchByName(ctx context.Context, firstName, lastName string) ([]*entity.User, error) {
	panic("")
}

func (s *UserQuery) SearchByIsDel(ctx context.Context, isDel bool) ([]*entity.User, error) {
	panic("")
}
