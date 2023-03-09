package operation

import (
	routing "github.com/gly-hub/fasthttp-routing"
	auth2 "github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/rpcserver"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/server/http"
)

// @title go-admin API
// @version 1.0.0
// @description 基于RpcX + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description
// @securityDefinitions.apikey Bearer
// @in header
// @name Operation

type OperationController struct {
	http.HttpController
}

// Login
// @Summary 登录
// @Description 用户登录
// @Tags 基础模块|用户登录登出
// @Param deptName body auth.LoginParams true "登录参数"
// @Success 200 {object} auth.LoginResp "{"code": 200, "data": [...]}"
// @Router /api/login [post]
func (a *OperationController) Login(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.Login, new(auth2.LoginParams), new(auth2.LoginResp))
}
