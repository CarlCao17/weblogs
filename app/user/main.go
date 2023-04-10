package main

import (
	"log"

	user "github.com/weblogs/kitex_gen/github/com/weblogs/user/userservice"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
