package rpc

import (
	"github.com/weblogs/kitex_gen/github/com/weblogs/blog/blogpostservice"
	"github.com/weblogs/kitex_gen/github/com/weblogs/user/userservice"
)

var (
	PostCli blogpostservice.Client
	UserCli userservice.Client
)

func InitRPC() {
	InitBlogRPC()
	InitUserRPC()
}

func InitBlogRPC() {
	PostCli = blogpostservice.MustNewClient()
}

func InitUserRPC() {
	UserCli = userservice.MustNewClient()
}
