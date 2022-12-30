package enum

import error_support "github.com/gly-hub/go-dandelion/error-support"

// 基础错误码 10000 - 11000

var (
	DataBaseErr   = &error_support.Error{Code: 11000, Msg: "数据库错误"}
	DataFormatErr = &error_support.Error{Code: 11001, Msg: "参数错误"}
)

// 功能错误码 11000 -

var (
	LoginPasswordErr      = &error_support.Error{Code: 12000, Msg: "用户密码错误"}
	LoginUserNameNotFound = &error_support.Error{Code: 12000, Msg: "用户不存在"}
)
