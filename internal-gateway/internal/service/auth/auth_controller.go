package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/server/http"
)

type AuthController struct {
	http.HttpController
}

// Login
// @Summary 登录
// @Description 用户登录
// @Tags 用户登录登出
// @Param deptName body auth.LoginParams true "登录参数"
// @Success 200 {object} auth.LoginResp "{"code": 200, "data": [...]}"
// @Router /api/login [post]
func (a *AuthController) Login(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.Login, new(auth.LoginParams), new(auth.LoginResp))
}

// Logout
// @Summary 登出
// @Description 用户登录
// @Tags 用户登录登出
// @Success 200 {object} string  "{"code": 200, "data": [...]}"
// @Router /api/logout [post]
func (a *AuthController) Logout(c *routing.Context) error {
	return nil
}
