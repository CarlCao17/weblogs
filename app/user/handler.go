package main

import (
	"context"

	"github.com/weblogs/app/user/domain/service"
	"github.com/weblogs/app/user/model/convert"
	"github.com/weblogs/kitex_gen/github/com/weblogs/user"
	"github.com/weblogs/pkg/errors"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = user.NewCreateUserResponse()
	if err = req.IsValid(); err != nil {
		return resp, errors.ParamErr.WithCause(err)
	}
	e, err := convert.CreateUserReq2Entity(req)
	if err != nil {
		return resp, errors.ParamErr.WithCause(err)
	}
	updateServiceI := service.GetUserUpdateServiceInstance()
	err = updateServiceI.CreateUser(ctx, e)
	if err != nil {
		return resp, err
	}
	resp.UserID = e.UserID
	resp.Role = e.UserRole.String()
	return
}

// DeleteUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (resp *user.DeleteUserResponse, err error) {
	resp = user.NewDeleteUserResponse()
	if err = req.IsValid(); err != nil {
		return resp, errors.ParamErr.WithCause(err)
	}
	e, err := convert.DeleteUserReq2Entity(req)
	if err != nil {
		return resp, errors.ParamErr.WithCause(err)
	}
	err = service.GetUserUpdateServiceInstance().DeleteUser(ctx, e)
	return
}

// SignIn implements the UserServiceImpl interface.
func (s *UserServiceImpl) SignIn(ctx context.Context, req *user.SignInRequest) (resp *user.SignInResponse, err error) {
	resp = user.NewSignInResponse()
	userID, userRole, err := service.GetUserQueryServiceInstance().Check(ctx, req.UserName, req.Password)
	if err != nil {
		return resp, err
	}
	resp.UserID = userID
	return resp, nil
}

// SignOut implements the UserServiceImpl interface.
func (s *UserServiceImpl) SignOut(ctx context.Context, req *user.SignOutRequest) (resp *user.SignOutResponse, err error) {
	// TODO: Your code here...
	return
}
