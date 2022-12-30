package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/application"
)

func (a *AuthController) GetMenuTree(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.UserMenu, new(auth.UserMenuTreeParams), new(auth.UserMenuTreeResp))
}
