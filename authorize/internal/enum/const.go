package enum

import error_support "github.com/gly-hub/go-dandelion/error-support"

var (
	DataBaseErr = &error_support.Error{Code: 100, Msg:  "数据库错误"}
)
