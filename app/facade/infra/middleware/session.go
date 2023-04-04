package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/weblogs/app/facade/model"
	"github.com/weblogs/pkg/errors"
	"github.com/weblogs/pkg/session"
)

const (
	CookieKey = "X-Access-Token"
)

func CookieMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		cookie := string(c.GetHeader(CookieKey))
		if cookie == "" {
			err := errors.CookieErr.WithMessage("unauthorized: have no access token")
			c.JSON(consts.StatusForbidden, model.Response{
				Code:    err.Code(),
				Message: err.Message(),
			})
			return
		}
		isLoggedIn := session.GetSessionInstance().IsLoggedIn(cookie)
		if !isLoggedIn {
			err := errors.CookieErr.WithMessage("unauthorized: invalid access token")
			c.JSON(consts.StatusForbidden, model.Response{
				Code:    err.Code(),
				Message: err.Message(),
			})
			return
		}
		c.Next(ctx)
	}
}
