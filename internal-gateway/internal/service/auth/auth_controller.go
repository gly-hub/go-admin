package auth

import (
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/server/http"
	routing "github.com/qiangxue/fasthttp-routing"
)

type AuthController struct {
	http.HttpController
}

func (a *AuthController) Login(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.Login, new(auth.LoginParams), new(auth.LoginResp))
}

func (a *AuthController) Logout(c *routing.Context) error {
	return nil
}
