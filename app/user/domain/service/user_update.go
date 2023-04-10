package service

import (
	"context"

	"github.com/weblogs/app/user/domain/repository"
	"github.com/weblogs/app/user/model/entity"
)

type UserUpdate struct {
	repo repository.UserRepo
}

func GetUserUpdateServiceInstance() *UserUpdate {
	if userUpdateService == nil {
		userUpdateService = NewUserUpdateService(repository.GetUserRepoInstance())
	}
	return userUpdateService
}

func NewUserUpdateService(repo repository.UserRepo) *UserUpdate {
	return &UserUpdate{
		repo: repo,
	}
}

var userUpdateService *UserUpdate

func (s *UserUpdate) CreateUser(ctx context.Context, e *entity.User) error {
	panic("")
}

func (s *UserUpdate) DeleteUser(ctx context.Context, e *entity.User) error {
	panic("")
}

func (s *UserUpdate) UpdateUser(ctx context.Context, e *entity.User) error {
	panic("")
}
