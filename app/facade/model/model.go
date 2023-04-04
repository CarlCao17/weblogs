package model

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/weblogs/pkg/errors"
)

type Response struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SendResponse(c *app.RequestContext, err error, data any) {
	ce := errors.ConvertError(err)
	c.JSON(http.StatusOK, Response{
		Code:    ce.Code(),
		Message: ce.Message(),
		Data:    data,
	})
}

type CreateBlogRequest struct {
	Title    string
	AuthorID int64
	Content  string
	Category *string
	Tags     []string
}

type UpdateBlogRequest struct {
	BlogID int64
	CreateBlogRequest
}

type DeleteBlogRequest struct {
	BlogID int64
}

type SignInRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
