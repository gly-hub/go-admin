package service

import (
	"context"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/logic"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/common"
	"github.com/gly-hub/go-dandelion/application"
	serr "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/tools/copyx"
	"strconv"
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
	return nil
}

// CreateAdminMenu 创建菜单项
func (ar *AuthRpc) CreateAdminMenu(ctx context.Context, req auth.AdminMenu, resp *common.Response) error {
	operator := application.GetHeader(ctx, "user_name")
	if err := logic.AdminMenu.CreateMenu(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// UpdateAdminMenu 更新菜单项
func (ar *AuthRpc) UpdateAdminMenu(ctx context.Context, req auth.AdminMenu, resp *common.Response) error {
	operator := application.GetHeader(ctx, "user_name")
	if err := logic.AdminMenu.UpdateMenu(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// DeleteAdminMenu 删除菜单项
func (ar *AuthRpc) DeleteAdminMenu(ctx context.Context, req auth.AdminMenu, resp *common.Response) error {
	operator := application.GetHeader(ctx, "user_name")
	if err := logic.AdminMenu.DeleteMenu(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// UserMenu 获取用户权限树
func (ar *AuthRpc) UserMenu(ctx context.Context, req auth.UserMenuTreeParams, resp *auth.UserMenuTreeResp) error {
	userId := application.GetHeader(ctx, "user_id")
	platform, err := strconv.Atoi(application.GetHeader(ctx, "platform"))
	if err != nil {
		serr.Format(enum.DataFormatErr, resp)
		return nil
	}
	menus, mErr := logic.AdminMenu.GetMenuTreeByUserId(userId, platform)
	if mErr != nil {
		serr.Format(mErr, resp)
		return nil
	}
	var respMenus []auth.AdminMenuTree
	_ = copyx.DeepCopy(&respMenus, &menus)
	resp.Menus = respMenus
	return nil
}
