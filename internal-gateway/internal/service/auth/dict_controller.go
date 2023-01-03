package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/common"
	"github.com/gly-hub/go-dandelion/application"
	"strconv"
)

func (a *AuthController) GetDictList(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.SearchAdminDictLabel, new(auth.SearchAdminDictParams), new(auth.SearchAdminDictResp))
}

func (a *AuthController) CreateDict(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.CreateAdminDict, new(auth.AdminDictLabel), new(common.Response))
}

func (a *AuthController) UpdateDict(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.UpdateAdminDict, new(auth.AdminDictLabel), new(common.Response))
}

func (a *AuthController) DeleteDict(c *routing.Context) error {
	var (
		req  = new(auth.AdminDictLabel)
		resp = new(common.Response)
	)
	req.Id, _ = strconv.Atoi(c.Param("id"))
	err := application.RpcCall(c, auth.ServerName, auth.DeleteAdminDict, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}

func (a *AuthController) GetDictListValue(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.SearchAdminDictValue, new(auth.SearchAdminValueParams), new(auth.SearchAdminValueResp))
}

func (a *AuthController) CreateDictValue(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.CreateAdminDictValue, new(auth.AdminDictValue), new(common.Response))
}

func (a *AuthController) UpdateDictValue(c *routing.Context) error {
	return application.SRpcCall(c, auth.ServerName, auth.UpdateAdminDictValue, new(auth.AdminDictValue), new(common.Response))
}

func (a *AuthController) DeleteDictValue(c *routing.Context) error {
	var (
		req  = new(auth.AdminDictValue)
		resp = new(common.Response)
	)
	req.Id, _ = strconv.Atoi(c.Param("id"))
	err := application.RpcCall(c, auth.ServerName, auth.DeleteAdminDictValue, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}
