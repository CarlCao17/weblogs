package conf

import (
	"os"
	"time"
)

const (
	EtcdAddress = "127.0.0.1:2379"

	BlogRpcServiceName   = "github.com.weblogs.blog"
	UserRpcServiceName   = "github.com.weblogs.user"
	BlogServiceAddress   = "127.0.0.1:8889"
	UserServiceAddress   = "127.0.0.1:8890"
	FacadeServiceAddress = "127.0.0.1:8080"
)

const (
	DataBaseName       = ""
	BlogCollectionName = "blog_post"
	BlogTableName      = "t_blog"
	UserTableName      = "t_user"
)

const (
	RPCTimeout        = 3 * time.Second
	RPCConnectTimeout = 50 * time.Millisecond
)

func DBName() string {
	return os.Getenv("DB_NAME")
}
