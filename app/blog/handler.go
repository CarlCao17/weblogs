package main

import (
	"context"

	domainservice "github.com/weblogs/app/blog/domain/service"
	"github.com/weblogs/app/blog/model/convert"
	"github.com/weblogs/kitex_gen/github/com/weblogs/blog"
)

// BlogPostServiceImpl implements the last service interface defined in the IDL.
type BlogPostServiceImpl struct{}

// CreatePost implements the BlogPostServiceImpl interface.
func (s *BlogPostServiceImpl) CreatePost(ctx context.Context, req *blog.CreatePostRequest) (resp *blog.CreatePostResponse, err error) {
	resp = &blog.CreatePostResponse{}
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	entity, err := convert.CreatePostReq2Entity(req)
	if err != nil {
		return resp, err
	}

	updateService := domainservice.GetBlogUpdateServiceInstance()
	err = updateService.CreatePost(ctx, entity)
	if err != nil {
		return resp, err
	}
	resp.BlogID = entity.PostID
	return
}

// UpdatePost implements the BlogPostServiceImpl interface.
func (s *BlogPostServiceImpl) UpdatePost(ctx context.Context, req *blog.UpdatePostRequest) (resp *blog.UpdatePostResponse, err error) {
	resp = &blog.UpdatePostResponse{}
	if err = req.IsValid(); err != nil {
		return resp, err
	}

	entity, err := convert.UpdatePostReq2Entity(req)
	if err != nil {
		return resp, err
	}

	updateService := domainservice.GetBlogUpdateServiceInstance()
	err = updateService.UpdatePost(ctx, entity)
	if err != nil {
		return resp, err
	}
	return
}

// DeletePost implements the BlogPostServiceImpl interface.
func (s *BlogPostServiceImpl) DeletePost(ctx context.Context, req *blog.DeletePostRequest) (resp *blog.DeletePostResponse, err error) {
	// TODO: Your code here...
	return
}
