package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/external-gateway/internal/enum"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/server/http"
)

type AuthController struct {
	http.HttpController
}

func (a *AuthController) Login(c *routing.Context) error {
	err := enum.DataBaseErr
	if err != nil {
		logger.Error(err)
		return a.Fail(c, err)
	}

	return a.Success(c, "123", "登录成功")
}
