package errors

import "github.com/weblogs/kitex_gen/github/com/weblogs/base"

func NewBaseRespWith(err error) *base.BaseResp {
	resp := NewSuccessResp()
	if err != nil {
		bizErr := ConvertError(err)
		resp.StatusCode = bizErr.Code()
		resp.StatusMessage = bizErr.Message()
	}
	return resp
}

func NewSuccessResp() *base.BaseResp {
	return &base.BaseResp{StatusCode: 0, StatusMessage: "success"}
}
