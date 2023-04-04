package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"

	"github.com/weblogs/app/facade/handler/blog"
	"github.com/weblogs/app/facade/infra/rpc"
)

func Init() {
	rpc.InitRPC()
}

func main() {
	Init()
	h := server.Default()

	// user service
	//userGroup := h.Group("/user")
	//userGroup.POST("/register", user.UserRegister)
	//userGroup.POST("/login", user.UserLogin)

	// blog service
	blogGroup := h.Group("/blog")
	blogGroup.POST("/create", blog.CreateBlog)
	h.Spin()
}
