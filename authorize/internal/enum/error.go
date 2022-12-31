package enum

import serr "github.com/gly-hub/go-dandelion/error-support"

// 基础错误码 10000 - 11000

var (
	DataBaseErr   = &serr.Error{Code: 11000, Msg: "数据库错误"}
	DataFormatErr = &serr.Error{Code: 11001, Msg: "参数错误"}
	SystemErr     = &serr.Error{Code: 11002, Msg: "系统错误"}
)

// 功能错误码 11000 -

var (
	LoginPasswordErr       = &serr.Error{Code: 12000, Msg: "用户密码错误"}
	LoginUserNameNotFound  = &serr.Error{Code: 12001, Msg: "用户不存在"}
	MenuNeedModule         = &serr.Error{Code: 12002, Msg: "一级菜单项必须为模块"}
	MenuParentMenuNotFound = &serr.Error{Code: 12003, Msg: "父级菜单项不存在"}
	PrimaryKeyNotFound     = &serr.Error{Code: 12004, Msg: "主键错误"}
	MenuNotFound           = &serr.Error{Code: 12005, Msg: "菜单不存在"}
)
