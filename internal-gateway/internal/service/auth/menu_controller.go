package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	auth2 "github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/lib"
	"github.com/gly-hub/go-admin/common/rpcserver"
	"github.com/gly-hub/go-dandelion/application"
	"strconv"
)

// GetMenuList
// @Summary 获取菜单树列表
// @Description 获取菜单树列表
// @Tags 权限模块|菜单管理
// @Param data body auth.SearchAdminMenuParams true "查询参数"
// @Success 200 {object} auth.UserMenuTreeResp "{"code": 200, "data": [...]}"
// @Router /api/auth/menu/search [post]
func (a *AuthController) GetMenuList(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.SearchAdminMenu, new(auth2.SearchAdminMenuParams), new(auth2.UserMenuTreeResp))
}

// CreateMenu
// @Summary 创建菜单项
// @Description 创建菜单项
// @Tags 权限模块|菜单管理
// @Param data body auth.AdminMenu true "菜单项参数"
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/menu [post]
func (a *AuthController) CreateMenu(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.CreateAdminMenu, new(auth2.AdminMenu), new(lib.Response))
}

// UpdateMenu
// @Summary 更新菜单项
// @Description 更新菜单项
// @Tags 权限模块|菜单管理
// @Param data body auth.AdminMenu true "菜单项参数"
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/menu [put]
func (a *AuthController) UpdateMenu(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.UpdateAdminMenu, new(auth2.AdminMenu), new(lib.Response))
}

// DeleteMenu
// @Summary 删除菜单项
// @Description 删除菜单项
// @Tags 权限模块|菜单管理
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/menu/:id [Delete]
func (a *AuthController) DeleteMenu(c *routing.Context) error {
	var (
		req  = new(auth2.AdminMenu)
		resp = new(lib.Response)
	)
	req.Id, _ = strconv.Atoi(c.Param("id"))
	err := application.RpcCall(c, rpcserver.ServerName, rpcserver.DeleteAdminMenu, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, resp.Msg)
}
