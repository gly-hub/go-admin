package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	auth2 "github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/lib"
	"github.com/gly-hub/go-admin/common/rpcserver"
	"github.com/gly-hub/go-dandelion/application"
	"strconv"
)

// GetDictList
// @Summary 获取字典标签列表
// @Description 获取字典标签列表
// @Tags 权限模块|字典管理
// @Param data body auth.SearchAdminDictParams true "查询参数"
// @Success 200 {object} auth.SearchAdminDictResp "{"code": 200, "data": [...]}"
// @Router /api/auth/dict/search [post]
func (a *AuthController) GetDictList(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.SearchAdminDictLabel, new(auth2.SearchAdminDictParams), new(auth2.SearchAdminDictResp))
}

// CreateDict
// @Summary 创建字典标签
// @Description 创建字典标签
// @Tags 权限模块|字典管理
// @Param data body auth.AdminDictLabel true "查询参数"
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/dict [post]
func (a *AuthController) CreateDict(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.CreateAdminDict, new(auth2.AdminDictLabel), new(lib.Response))
}

// UpdateDict
// @Summary 更新字典标签
// @Description 更新字典标签
// @Tags 权限模块|字典管理
// @Param data body auth.AdminDictLabel true "查询参数"
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/dict [put]
func (a *AuthController) UpdateDict(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.UpdateAdminDict, new(auth2.AdminDictLabel), new(lib.Response))
}

// DeleteDict
// @Summary 删除字典标签
// @Description 删除字典标签
// @Tags 权限模块|字典管理
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/dict/:id [Delete]
func (a *AuthController) DeleteDict(c *routing.Context) error {
	var (
		req  = new(auth2.AdminDictLabel)
		resp = new(lib.Response)
	)
	req.Id, _ = strconv.Atoi(c.Param("id"))
	err := application.RpcCall(c, rpcserver.ServerName, rpcserver.DeleteAdminDict, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}

// GetDictValueList
// @Summary 获取字典值列表
// @Description 获取字典值列表
// @Tags 权限模块|字典管理
// @Param data body auth.SearchAdminValueParams true "查询参数"
// @Success 200 {object} auth.SearchAdminValueResp "{"code": 200, "data": [...]}"
// @Router /api/auth/dict/value/search [post]
func (a *AuthController) GetDictValueList(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.SearchAdminDictValue, new(auth2.SearchAdminValueParams), new(auth2.SearchAdminValueResp))
}

// CreateDictValue
// @Summary 创建字典值
// @Description 创建字典值
// @Tags 权限模块|字典管理
// @Param data body auth.AdminDictLabel true "查询参数"
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/dict/value [post]
func (a *AuthController) CreateDictValue(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.CreateAdminDictValue, new(auth2.AdminDictValue), new(lib.Response))
}

// UpdateDictValue
// @Summary 更新字典值
// @Description 更新字典值
// @Tags 权限模块|字典管理
// @Param data body auth.AdminDictLabel true "查询参数"
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/dict/value [put]
func (a *AuthController) UpdateDictValue(c *routing.Context) error {
	return application.SRpcCall(c, rpcserver.ServerName, rpcserver.UpdateAdminDictValue, new(auth2.AdminDictValue), new(lib.Response))
}

// DeleteDictValue
// @Summary 删除字典值
// @Description 删除字典值
// @Tags 权限模块|字典管理
// @Success 200 {object} lib.Response "{"code": 200, "data": [...]}"
// @Router /api/auth/dict/value/:id [Delete]
func (a *AuthController) DeleteDictValue(c *routing.Context) error {
	var (
		req  = new(auth2.AdminDictValue)
		resp = new(lib.Response)
	)
	req.Id, _ = strconv.Atoi(c.Param("id"))
	err := application.RpcCall(c, rpcserver.ServerName, rpcserver.DeleteAdminDictValue, req, resp)
	if err != nil {
		return a.Fail(c, err)
	}
	return a.Success(c, resp, "")
}
