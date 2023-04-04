// Code generated by Kitex v0.5.0. DO NOT EDIT.

package blogpostservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	blog "github.com/weblogs/kitex_gen/github/com/weblogs/blog"
)

func serviceInfo() *kitex.ServiceInfo {
	return blogPostServiceServiceInfo
}

var blogPostServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "BlogPostService"
	handlerType := (*blog.BlogPostService)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreatePost": kitex.NewMethodInfo(createPostHandler, newBlogPostServiceCreatePostArgs, newBlogPostServiceCreatePostResult, false),
		"UpdatePost": kitex.NewMethodInfo(updatePostHandler, newBlogPostServiceUpdatePostArgs, newBlogPostServiceUpdatePostResult, false),
		"DeletePost": kitex.NewMethodInfo(deletePostHandler, newBlogPostServiceDeletePostArgs, newBlogPostServiceDeletePostResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "blog",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.0",
		Extra:           extra,
	}
	return svcInfo
}

func createPostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*blog.BlogPostServiceCreatePostArgs)
	realResult := result.(*blog.BlogPostServiceCreatePostResult)
	success, err := handler.(blog.BlogPostService).CreatePost(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBlogPostServiceCreatePostArgs() interface{} {
	return blog.NewBlogPostServiceCreatePostArgs()
}

func newBlogPostServiceCreatePostResult() interface{} {
	return blog.NewBlogPostServiceCreatePostResult()
}

func updatePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*blog.BlogPostServiceUpdatePostArgs)
	realResult := result.(*blog.BlogPostServiceUpdatePostResult)
	success, err := handler.(blog.BlogPostService).UpdatePost(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBlogPostServiceUpdatePostArgs() interface{} {
	return blog.NewBlogPostServiceUpdatePostArgs()
}

func newBlogPostServiceUpdatePostResult() interface{} {
	return blog.NewBlogPostServiceUpdatePostResult()
}

func deletePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*blog.BlogPostServiceDeletePostArgs)
	realResult := result.(*blog.BlogPostServiceDeletePostResult)
	success, err := handler.(blog.BlogPostService).DeletePost(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newBlogPostServiceDeletePostArgs() interface{} {
	return blog.NewBlogPostServiceDeletePostArgs()
}

func newBlogPostServiceDeletePostResult() interface{} {
	return blog.NewBlogPostServiceDeletePostResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreatePost(ctx context.Context, req *blog.CreatePostRequest) (r *blog.CreatePostResponse, err error) {
	var _args blog.BlogPostServiceCreatePostArgs
	_args.Req = req
	var _result blog.BlogPostServiceCreatePostResult
	if err = p.c.Call(ctx, "CreatePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdatePost(ctx context.Context, req *blog.UpdatePostRequest) (r *blog.UpdatePostResponse, err error) {
	var _args blog.BlogPostServiceUpdatePostArgs
	_args.Req = req
	var _result blog.BlogPostServiceUpdatePostResult
	if err = p.c.Call(ctx, "UpdatePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePost(ctx context.Context, req *blog.DeletePostRequest) (r *blog.DeletePostResponse, err error) {
	var _args blog.BlogPostServiceDeletePostArgs
	_args.Req = req
	var _result blog.BlogPostServiceDeletePostResult
	if err = p.c.Call(ctx, "DeletePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
