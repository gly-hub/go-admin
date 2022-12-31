package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/common"
	"github.com/gly-hub/go-dandelion/application"
	"strconv"
)

func (a *AuthController) GetMenuTree(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.UserMenu, new(auth.UserMenuTreeParams), new(auth.UserMenuTreeResp))
}

func (a *AuthController) GetMenuList(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.SearchAdminMenu, new(auth.SearchAdminMenuParams), new(auth.UserMenuTreeResp))
}

func (a *AuthController) CreateMenu(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.CreateAdminMenu, new(auth.AdminMenu), new(common.Response))
}

func (a *AuthController) UpdateMenu(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.UpdateAdminMenu, new(auth.AdminMenu), new(common.Response))
}

func (a *AuthController) DeleteMenu(c *routing.Context) error {
	var (
		req  = new(auth.AdminMenu)
		resp = new(common.Response)
	)
	req.Id, _ = strconv.Atoi(c.Param("id"))
	err := application.RpcCall(c, auth.ServerName, auth.DeleteAdminMenu, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}
