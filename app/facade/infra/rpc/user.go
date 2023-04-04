package rpc

import (
	"context"
	"sync"

	"github.com/weblogs/kitex_gen/github/com/weblogs/user"
	"github.com/weblogs/kitex_gen/github/com/weblogs/user/userservice"
	"github.com/weblogs/pkg/errors"
)

type User interface {
	CreateUser(ctx context.Context, userName, password, email, firstName, lastName string) (int64, error)
	SignIn(ctx context.Context, userName, password string) (int64, error)
}

type userImpl struct {
	cli userservice.Client
}

func NewUser(cli userservice.Client) User {
	return userImpl{
		cli: cli,
	}
}

var (
	userI User
	once  sync.Once
)

func GetUserInstance() User {
	if userI == nil {
		once.Do(func() {
			userI = NewUser(UserCli)
		})
	}
	return userI
}

func (u userImpl) CreateUser(ctx context.Context, userName, password, email, firstName, lastName string) (int64, error) {
	req := &user.CreateUserRequest{
		UserName:  userName,
		Password:  password,
		Email:     email,
		FirstName: &firstName,
		LastName:  &lastName,
	}

	resp, err := u.cli.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp == nil || resp.BaseResp == nil {
		return 0, errors.RPCNilResponseErr
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.NewBizError(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.UserID, nil
}

func (u userImpl) SignIn(ctx context.Context, userName, password string) (int64, error) {
	req := &user.LoginRequest{
		UserName: userName,
		Password: password,
	}

	resp, err := u.cli.Login(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp == nil || resp.BaseResp == nil {
		return 0, errors.RPCNilResponseErr
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errors.NewBizError(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.UserID, nil
}
