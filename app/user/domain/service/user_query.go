package service

import (
	"context"

	"github.com/cockroachdb/errors"
	"github.com/weblogs/app/user/model/po"
	"github.com/weblogs/pkg/session"

	"github.com/weblogs/app/user/domain/repository"
	"github.com/weblogs/app/user/domain/service/util"
	"github.com/weblogs/app/user/model/entity"
	"github.com/weblogs/app/user/model/po/field"
)

type UserQuery struct {
	repo repository.UserRepo
}

func GetUserQueryServiceInstance() *UserQuery {
	if userQueryService == nil {
		userQueryService = NewUserQueryService(repository.GetUserRepoInstance())
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

}

func (s *UserQuery) MGet(ctx context.Context, userIDs []int64) ([]*entity.User, error) {

}

func (s *UserQuery) Check(ctx context.Context, userName, passwd string) (int64, po.UserRole, error) {
	encryptedPasswd := util.EncryptPasswd([]byte(passwd))
	filters := map[field.Field]any{
		field.UserName: userName,
		field.Password: encryptedPasswd,
	}
	users, err := s.repo.Query(ctx, filters, field.UserID, field.UserRole)
	if err != nil {
		return 0, po.UserRole(0), errors.Wrapf(err, "userquery: check username and password failed, username=%s", userName)
	}
	if len(users) == 0 {
		return 0, po.UserRole(0), errors.Newf("userquery: could not find the record with user=%s, passwd=%s", userName, passwd)
	}
	userID, userRole := users[0].UserID, users[0].UserRole
	session.GetSessionInstance().Set(session.Model{UserID: })
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
