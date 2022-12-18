package auth

import "github.com/gly-hub/go-admin/common/model/common"

type LoginParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResp struct {
	common.Response
	Token string `json:"token"`
}

type LogoutParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LogoutResp struct {
	common.Response
}
