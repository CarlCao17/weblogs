package session

import (
	"context"
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/go-session/session/v3"
)

type ErrorHandleFunc func(ctx context.Context, c *app.RequestContext, err error)

type Config struct {
	ErrorHandleFunc ErrorHandleFunc
	StoreKey        string
	ManagerKey      string
}

var (
	managerKey string
	storeKey   string

	DefaultConfig = Config{
		ErrorHandleFunc: func(ctx context.Context, c *app.RequestContext, err error) {
			c.AbortWithError(500, err)
		},
		StoreKey:   "Key-Session-Store",
		ManagerKey: "Key-Session-Manager",
	}
)

func New(opt ...session.Option) app.HandlerFunc {
	return NewWithConfig(DefaultConfig, opt...)
}

// TODO: 改造 go-session 否则无法直接使用
func NewWithConfig(config Config, opt ...session.Option) app.HandlerFunc {
	if config.ErrorHandleFunc == nil {
		config.ErrorHandleFunc = DefaultConfig.ErrorHandleFunc
	}

	if config.ManagerKey == "" {
		config.ManagerKey = DefaultConfig.ManagerKey
	}
	managerKey = config.ManagerKey

	if config.StoreKey == "" {
		config.StoreKey = DefaultConfig.StoreKey
	}
	storeKey = config.StoreKey

	manager := session.NewManager(opt...)
	return func(ctx context.Context, c *app.RequestContext) {
		c.Set(managerKey, manager)

		req, _ := adaptor.GetCompatRequest(&c.Request)
		resp := adaptor.GetCompatResponseWriter(&c.Response)

		store, err := manager.Start(nil, resp, req)
		if err != nil {
			config.ErrorHandleFunc(ctx, c, err)
			return
		}
		c.Set(storeKey, store)
		c.Next(ctx)
	}
}

func FromContext(ctx *app.RequestContext) (store session.Store, exists bool) {
	storeAny, exists := ctx.Get(storeKey)
	if !exists {
		return nil, exists
	}
	store = storeAny.(session.Store)
	return store, exists
}

var (
	ErrManagerNotExists = fmt.Errorf("session manager does not exist in context")
)

func Destroy(ctx *app.RequestContext) error {
	managerAny, exists := ctx.Get(managerKey)
	if !exists {
		return ErrManagerNotExists
	}
	manager := managerAny.(*session.Manager)

	req, _ := adaptor.GetCompatRequest(&ctx.Request)
	resp := adaptor.GetCompatResponseWriter(&ctx.Response)
	return manager.Destroy(nil, resp, req)
}

func Refresh(ctx *app.RequestContext) (session.Store, error) {
	managerAny, exists := ctx.Get(managerKey)
	if !exists {
		return nil, ErrManagerNotExists
	}
	manager := managerAny.(*session.Manager)
	req, _ := adaptor.GetCompatRequest(&ctx.Request)
	resp := adaptor.GetCompatResponseWriter(&ctx.Response)
	return manager.Refresh(nil, resp, req)
}
