package service

import (
	"context"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/logic"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/application"
	serr "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/tools/copyx"
	"strconv"
)

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
