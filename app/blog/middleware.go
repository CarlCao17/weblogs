package main

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"

	"github.com/weblogs/kitex_gen/github/com/weblogs/base"
	"github.com/weblogs/pkg"
	"github.com/weblogs/pkg/errors"
	"github.com/weblogs/pkg/log"
	"github.com/weblogs/pkg/util"
)

func ReqRespMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		rpcInfo := rpcinfo.GetRPCInfo(ctx)
		rpcCtx := pkg.RPCContext{
			Context: ctx,
			Start:   time.Now().UTC(),
			Method:  rpcInfo.To().Method(),
		}
		// 1. 打印请求日志
		log.CtxInfof(rpcCtx, "[RPCRequest] method=%s, req=%v", rpcCtx.Method, req)
		defer func() {
			rpcCtx.MileStone = time.Now().UTC()
			if err != nil {
				log.CtxErrorf(rpcCtx, "rpc return error, method=%s, err=%v", rpcCtx.Method, err)
			}
			// 2. Recover from panic and transfer the panic to error
			if r := recover(); r != nil {
				log.CtxErrorf(rpcCtx, "rpc panicked, method=%s, panic=%v", rpcCtx.Method, r)
				if err == nil {
					err = errors.ServiceErr.WithMessage(fmt.Sprintf("recover from panic=%v", r))
				}
			}
			// 1. 打印请求日志和耗时
			log.CtxInfof(rpcCtx, "[RPCResponse] method=%s, cost=%v, resp=%v", rpcCtx.Method, rpcCtx.MileStone.Sub(rpcCtx.Start), resp)

			// 3. 处理 BaseResp，根据 error 填充，不返回业务 error 给框架
			gresp, ok := resp.(util.GenericResponse)
			if !ok {
				log.CtxErrorf(rpcCtx, "rpc returned error have no SetBaseResp or GetBaseResp method, method=%s", rpcCtx.Method)
			}
			gresp.SetBaseResp(BuildBaseResp(err))
			err = nil
		}()
		err = next(rpcCtx, req, resp)
		return err
	}
}

func BuildBaseResp(err error) *base.BaseResp {
	if err == nil {
		return &base.BaseResp{StatusMessage: "success", StatusCode: 0}
	}
	bizErr := errors.ConvertError(err)
	return &base.BaseResp{StatusMessage: bizErr.Message(), StatusCode: bizErr.Code()}
}
