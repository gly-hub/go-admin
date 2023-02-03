package service

import (
	"context"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/logic"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/application"
	error_support "github.com/gly-hub/go-dandelion/error-support"
	serr "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/tools/copyx"
	"strconv"
)

// Login 登录
func (ar *AuthRpc) Login(ctx context.Context, req auth.LoginParams, resp *auth.LoginResp) error {
	var (
		token string
		err   error
	)
	ip := application.GetHeader(ctx, "ip")
	if token, err = logic.Auth.Login(req.UserId, req.Password, ip); err != nil {
		error_support.Format(err, resp)
		return nil
	}
	resp.Token = token
	return nil
}

// Logout 登出
func (ar *AuthRpc) Logout(ctx context.Context, req auth.LogoutParams, resp *auth.LogoutResp) error {
	userId := application.GetHeader(ctx, "userId")
	if err := logic.Auth.Logout(userId); err != nil {
		error_support.Format(err, resp)
		return nil
	}
	return nil
}

// CheckToken token校验
func (ar *AuthRpc) CheckToken(ctx context.Context, req auth.CheckTokenParams, resp *auth.CheckTokenResp) error {
	if meta, newToken, err := logic.Auth.CheckToken(req.Token); err != nil {
		error_support.Format(err, resp)
		return nil
	} else {
		_ = copyx.DeepCopy(&resp, &meta)
		resp.NewToken = newToken
	}
	return nil
}

// SystemInfo 获取系统信息
func (ar *AuthRpc) SystemInfo(ctx context.Context, req auth.SystemInfoSearch, resp *auth.SystemInfoResp) error {
	resp.AppLogo = "https://doc-image.zhangwj.com/img/go-admin.png"
	resp.AppName = "go-admin"
	return nil
}

// UserMenu 获取用户权限树
func (ar *AuthRpc) UserMenu(ctx context.Context, req auth.UserMenuTreeParams, resp *auth.UserMenuTreeResp) error {
	userId := application.GetHeader(ctx, "userId")
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
