package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/weblogs/app/facade/infra/rpc"
	"github.com/weblogs/app/facade/model"
	"github.com/weblogs/pkg/errors"
)

func SignIn(ctx context.Context, c *app.RequestContext) {
	var signInReq model.SignInRequest
	if err := c.BindAndValidate(&signInReq); err != nil {
		model.SendResponse(c, errors.ParamErr.WithCause(err), nil)
		return
	}
	userID, err := rpc.GetUserInstance().SignIn(ctx, signInReq.UserName, signInReq.Password)
	if err != nil {
		model.SendResponse(c, errors.SignInErr.WithCause(err), nil)
		return
	}
	model.SendResponse(c, nil, map[string]any{
		"user_id": userID,
	})
}

func SignUp(ctx context.Context, c *app.RequestContext) {
	var signUpReq model.SignUpRequest
	if err := c.BindAndValidate(&signUpReq); err != nil {
		model.SendResponse(c, errors.ParamErr.WithCause(err), nil)
		return
	}
	userID, err := rpc.GetUserInstance().SignUp(ctx, &signUpReq)
	if err != nil {
		model.SendResponse(c, errors.UserNotExistErr, nil)
		return
	}
	model.SendResponse(c, nil, map[string]any{
		"user_id": userID,
	})

}
