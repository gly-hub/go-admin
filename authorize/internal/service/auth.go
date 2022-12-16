package service

import (
	"context"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/logger"
)

func (ar *AuthRpc) Login(ctx context.Context, req auth.LoginParams, resp *auth.LoginResp)(err error){
	logger.Info(req.UserName, req.Password)
	logger.Debug("1234")
	resp.Data = "5432"
	return nil
}
