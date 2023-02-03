package auth

import "github.com/gly-hub/go-admin/common/model/lib"

type SystemInfoSearch struct {
}

type SystemInfoResp struct {
	lib.Response
	AppLogo string `json:"appLogo"`
	AppName string `json:"appName"`
}
