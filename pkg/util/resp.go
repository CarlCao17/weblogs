package util

import "github.com/weblogs/kitex_gen/github/com/weblogs/base"

type GenericResponse interface {
	GetBaseResp() *base.BaseResp
	SetBaseResp(br *base.BaseResp)
}
