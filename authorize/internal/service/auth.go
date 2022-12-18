package service

import (
	"context"
	"github.com/gly-hub/go-admin/authorize/internal/logic"
	"github.com/gly-hub/go-admin/common/model/auth"
	error_support "github.com/gly-hub/go-dandelion/error-support"
)

func (ar *AuthRpc) Login(ctx context.Context, req auth.LoginParams, resp *auth.LoginResp)error{
	var (
		token string
		err error
	)
	if token, err = logic.Auth.Login(req.UserName, req.Password); err != nil{
		error_support.Format(err, resp)
		return nil
	}
	resp.Token = token
	return nil
}

func (ar *AuthRpc) Logout(ctx context.Context, req auth.LogoutParams, resp *auth.LogoutResp)error{
	if err := logic.Auth.Logout(req.UserName); err != nil{
		error_support.Format(err, resp)
		return nil
	}
	return nil
}
