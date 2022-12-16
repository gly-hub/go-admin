package auth

import (
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/server/http"
	routing "github.com/qiangxue/fasthttp-routing"
)

type AuthController struct {
	http.HttpController
}

func (a *AuthController) Login(c *routing.Context) error {
	req := auth.LoginParams{
		UserName: "1234",
		Password: "2345",
	}

	logger.Debug(req)
	resp := &auth.LoginResp{}

	err := application.RpcCall(auth.ServerName, auth.Login, req, resp)
	if err != nil{
		logger.Error(err)
		return a.Fail(c, err)
	}

	return a.Success(c, resp.Data, "登录成功")
}
