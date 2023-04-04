package blog

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/spf13/cast"

	"github.com/weblogs/app/facade/infra/rpc"
	"github.com/weblogs/app/facade/model"
	"github.com/weblogs/kitex_gen/github/com/weblogs/blog"
	"github.com/weblogs/pkg/errors"
)

// CreateBlog godoc
// @Summary create blog
// @Description
// @Tags
// @Accept json
// @Product json
// @Param
// @Security
// @Success
// @Router /blog/create [post]
func CreateBlog(ctx context.Context, c *app.RequestContext) {
	var createReq model.CreateBlogRequest
	if err := c.BindAndValidate(&createReq); err != nil {
		model.SendResponse(c, errors.ParamErr.WithCause(err), nil)
		return
	}
	req := &blog.CreateBlogRequest{
		Title:    createReq.Title,
		AuthorID: createReq.AuthorID,
		Content:  createReq.Content,
		Category: createReq.Category,
		Tags:     createReq.Tags,
	}
	blogID, status, err := rpc.NewBlogRPC().CreateBlogRaw(ctx, req)
	if err != nil {
		model.SendResponse(c, errors.ConvertError(err), nil)
		return
	}
	model.SendResponse(c, errors.Success, map[string]any{
		"blog_id": cast.ToString(blogID),
		"status":  status,
	})
}
