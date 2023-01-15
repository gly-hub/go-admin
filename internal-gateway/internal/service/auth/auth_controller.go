package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	auth2 "github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/rpcserver"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/server/http"
)

type AuthController struct {
	http.HttpController
}

// Login
// @Summary 登录
// @Description 用户登录
// @Tags 基础模块|用户登录登出
// @Param deptName body auth.LoginParams true "登录参数"
// @Success 200 {object} auth.LoginResp "{"code": 200, "data": [...]}"
// @Router /api/login [post]
func (a *AuthController) Login(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.Login, new(auth2.LoginParams), new(auth2.LoginResp))
}

// Logout
// @Summary 登出
// @Description 用户登录
// @Tags 基础模块|用户登录登出
// @Success 200 {object} string  "{"code": 200, "data": [...]}"
// @Router /api/logout [post]
func (a *AuthController) Logout(c *routing.Context) error {
	return nil
}

// GetSystemInfo
// @Summary 获取系统信息
// @Description 获取系统信息
// @Tags 基础模块|基础路由
// @Success 200 {object} auth.SystemInfoResp "{"code": 200, "data": [...]}"
// @Router /api/auth/system/menu/tree [get]
func (a *AuthController) GetSystemInfo(c *routing.Context) error {
	var (
		req  = new(auth2.SystemInfoSearch)
		resp = new(auth2.SystemInfoResp)
	)
	err := application.RpcCall(c, rpcserver.ServerName, rpcserver.SystemInfo, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}

// GetMenuTree
// @Summary 获取侧边栏菜单树
// @Description 获取侧边栏菜单树
// @Tags 基础模块|基础路由
// @Success 200 {object} auth.UserMenuTreeResp "{"code": 200, "data": [...]}"
// @Router /api/auth/system/menu/tree [post]
func (a *AuthController) GetMenuTree(c *routing.Context) error {
	var (
		req  = new(auth2.UserMenuTreeParams)
		resp = new(auth2.UserMenuTreeResp)
	)
	err := application.RpcCall(c, rpcserver.ServerName, rpcserver.UserMenu, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}
