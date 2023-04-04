package errors

import (
	"errors"
	"fmt"
)

const (
	SuccessCode           = 0
	ServiceErrCode        = 10001
	ParamErrCode          = 10002
	RPCNilResponseErrCode = 10003

	// User Error Code
	LoginErrCode            = 11001
	UserNotExistErrCode     = 11002
	UserAlreadyExistErrCode = 11003
	CookieErrCode           = 11004
)

type BizError struct {
	code  int64
	msg   string
	cause error
}

func NewBizError(code int64, msg string) *BizError {
	return &BizError{
		code: code,
		msg:  msg,
	}
}

func (e *BizError) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s, cause=%v", e.code, e.msg, e.cause)
}

func (e *BizError) String() string {
	return fmt.Sprintf("%s: %s, cause=%s", codeStr(e.code), e.msg, e.cause)
}

func (e *BizError) WithMessage(msg string) *BizError {
	e.msg = msg
	return e
}

func (e *BizError) WithCause(c error) *BizError {
	e.cause = c
	return e
}

func (e *BizError) Code() int64 {
	return e.code
}

func (e *BizError) Message() string {
	return e.msg
}

func (e *BizError) Cause() error {
	return e.cause
}

var (
	Success             = NewBizError(SuccessCode, "success")
	ServiceErr          = NewBizError(ServiceErrCode, "service is unusable temporary")
	RPCNilResponseErr   = NewBizError(RPCNilResponseErrCode, "rpc response or BaseResp is nil")
	ParamErr            = NewBizError(ParamErrCode, "request parameter is wrong")
	SignInErr           = NewBizError(LoginErrCode, "username or password is wrong")
	UserNotExistErr     = NewBizError(UserNotExistErrCode, "user does not exist")
	UserAlreadyExistErr = NewBizError(UserAlreadyExistErrCode, "user has already exist")
	CookieErr           = NewBizError(CookieErrCode, "cookie error")
)

var (
	m = map[int64]string{
		ServiceErrCode:          "service internal error",
		ParamErrCode:            "param error",
		LoginErrCode:            "login error",
		UserNotExistErrCode:     "user not exist",
		UserAlreadyExistErrCode: "user already exist",
	}
)

func codeStr(code int64) string {
	return m[code]
}

func ConvertError(err error) *BizError {
	e := &BizError{}
	if errors.As(err, e) {
		return e
	}
	return ServiceErr.WithMessage(e.Error())
}
