package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/weblogs/app/facade/handler/blog"
	"github.com/weblogs/app/facade/handler/user"
	"github.com/weblogs/app/facade/infra/rpc"
	"github.com/weblogs/pkg/session"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	h := server.Default()

	h.Use(session.Middleware)

	// user service
	userGroup := h.Group("/user")
	userGroup.POST("/signup", user.SignUp)
	userGroup.POST("/signin", user.SignIn)

	// blog service
	blogGroup := h.Group("/blog")
	blogGroup.POST("/create", blog.CreateBlog)
	h.Spin()
}
