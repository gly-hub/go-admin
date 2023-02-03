package service

import (
	"context"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/logic"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/lib"
	"github.com/gly-hub/go-dandelion/application"
	serr "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/tools/copyx"
)

// SearchAdminMenu 查询菜单列表
func (ar *AuthRpc) SearchAdminMenu(ctx context.Context, req auth.SearchAdminMenuParams, resp *auth.UserMenuTreeResp) error {
	menus, err := logic.AdminMenu.GetMenuList(req)
	if err != nil {
		serr.Format(err, resp)
		return nil
	}

	var respMenus []auth.AdminMenuTree
	_ = copyx.DeepCopy(&respMenus, &menus)
	resp.Menus = respMenus
	resp.Msg = enum.RespSearchSuccess
	return nil
}

// CreateAdminMenu 创建菜单项
func (ar *AuthRpc) CreateAdminMenu(ctx context.Context, req auth.AdminMenu, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminMenu.CreateMenu(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	resp.Msg = enum.RespCreateSuccess
	return nil
}

// UpdateAdminMenu 更新菜单项
func (ar *AuthRpc) UpdateAdminMenu(ctx context.Context, req auth.AdminMenu, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminMenu.UpdateMenu(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	resp.Msg = enum.RespUpdateSuccess
	return nil
}

// DeleteAdminMenu 删除菜单项
func (ar *AuthRpc) DeleteAdminMenu(ctx context.Context, req auth.AdminMenu, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminMenu.DeleteMenu(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	resp.Msg = enum.RespDeleteSuccess
	return nil
}
