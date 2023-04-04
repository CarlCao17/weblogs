package main

import (
	"log"

	"github.com/cloudwego/kitex/server"

	blog "github.com/weblogs/kitex_gen/github/com/weblogs/blog/blogpostservice"
	"github.com/weblogs/pkg/mongo"
)

func main() {
	svr := blog.NewServer(new(BlogPostServiceImpl), server.WithMiddleware(ReqRespMiddleware))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}

func InitServer() {
	mongo.InitMongo()
}
