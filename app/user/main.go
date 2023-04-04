package main

import (
	user "github.com/weblogs/app/user/kitex_gen/github/com/weblogs/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
