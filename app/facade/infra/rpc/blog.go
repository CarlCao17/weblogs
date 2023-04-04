package rpc

import (
	"context"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"

	"github.com/weblogs/kitex_gen/github/com/weblogs/blog"
	"github.com/weblogs/kitex_gen/github/com/weblogs/blog/blogpostservice"
	"github.com/weblogs/kitex_gen/github/com/weblogs/model"
	"github.com/weblogs/pkg/conf"
	"github.com/weblogs/pkg/errors"
	"github.com/weblogs/pkg/log"
)

type Post interface {
	CreatePost(ctx context.Context, title string, authorID int64, content string) (int64, model.BlogStatus, error)
	CreatePostRaw(ctx context.Context, req *blog.CreatePostRequest) (int64, model.BlogStatus, error)
	UpdatePost(ctx context.Context, PostID, authorID int64, title, content, category *string, tags []string) (bool, error)
	DeletePost(ctx context.Context, PostID int64) (bool, error)
}

type PostRpcImpl struct {
	blogpostservice.Client
}

func (b PostRpcImpl) CreatePost(ctx context.Context, title string, authorID int64, content string) (int64, model.BlogStatus, error) {
	req := &blog.CreatePostRequest{
		Title:    title,
		AuthorID: authorID,
		Content:  content,
	}
	return b.CreatePostRaw(ctx, req)
}

func (b PostRpcImpl) CreatePostRaw(ctx context.Context, req *blog.CreatePostRequest) (int64, model.BlogStatus, error) {
	resp, err := b.Client.CreatePost(ctx, req)
	if err != nil {
		return 0, model.BlogStatus_NotExist, err
	}
	if resp == nil || resp.BaseResp == nil {
		return 0, model.BlogStatus_NotExist, errors.RPCNilResponseErr
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, model.BlogStatus_NotExist, errors.NewBizError(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMessage)
	}
	return resp.BlogID, resp.Status, nil
}

func (b PostRpcImpl) UpdatePost(ctx context.Context, PostID, authorID int64, title, content, category *string, tags []string) (bool, error) {
	req := &blog.UpdatePostRequest{
		BlogID:   PostID,
		AuthorID: authorID,
		Title:    title,
		Content:  content,
		Category: category,
		Tags:     tags,
	}
	resp, err := b.Client.UpdatePost(ctx, req)
	if err != nil {
		return false, err
	}
	if resp == nil || resp.BaseResp == nil {
		return false, errors.RPCNilResponseErr
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errors.NewBizError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return true, nil
}

func (b PostRpcImpl) DeletePost(ctx context.Context, PostID int64) (bool, error) {
	req := &blog.DeletePostRequest{
		BlogID: PostID,
	}
	resp, err := b.Client.DeletePost(ctx, req)
	if err != nil {
		return false, err
	}
	if resp == nil || resp.BaseResp == nil {
		return false, errors.RPCNilResponseErr
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errors.NewBizError(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return true, nil
}

var (
	_ = NewPostRPC
	_ = InitPostRPC
)

func NewPostRPC() Post {
	return PostRpcImpl{PostClient}
}

var PostClient blogpostservice.Client

func InitPostRPC() {
	r, err := etcd.NewEtcdResolver([]string{conf.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := blogpostservice.NewClient(
		conf.BlogRpcServiceName,
		client.WithRPCTimeout(conf.RPCTimeout),
		client.WithConnectTimeout(conf.RPCConnectTimeout),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	PostClient = c
	log.Infof("Init Post rpc success")
}
