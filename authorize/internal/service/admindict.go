package service

import (
	"context"
	"github.com/gly-hub/go-admin/authorize/internal/logic"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-admin/common/model/lib"
	"github.com/gly-hub/go-dandelion/application"
	serr "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/tools/copyx"
)

// SearchAdminDictLabel 查询字典列表
func (ar *AuthRpc) SearchAdminDictLabel(ctx context.Context, req auth.SearchAdminDictParams, resp *auth.SearchAdminDictResp) error {
	labels, total, err := logic.AdminDict.GetDictList(req)
	if err != nil {
		serr.Format(err, resp)
		return nil
	}

	var respLabels []auth.AdminDictLabel
	_ = copyx.DeepCopy(&respLabels, &labels)
	resp.List = respLabels
	resp.Total = total
	return nil
}

// CreateAdminDict 创建字典标签
func (ar *AuthRpc) CreateAdminDict(ctx context.Context, req auth.AdminDictLabel, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminDict.CreateDict(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// UpdateAdminDict 更新字典标签
func (ar *AuthRpc) UpdateAdminDict(ctx context.Context, req auth.AdminDictLabel, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminDict.UpdateDict(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// DeleteAdminDict 删除字典标签
func (ar *AuthRpc) DeleteAdminDict(ctx context.Context, req auth.AdminDictLabel, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminDict.DeleteDict(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// SearchAdminDictValue 查询字典值列表
func (ar *AuthRpc) SearchAdminDictValue(ctx context.Context, req auth.SearchAdminValueParams, resp *auth.SearchAdminValueResp) error {
	values, total, err := logic.AdminDict.GetDictValueList(req)
	if err != nil {
		serr.Format(err, resp)
		return nil
	}

	var respValues []auth.AdminDictValue
	_ = copyx.DeepCopy(&respValues, &values)
	resp.List = respValues
	resp.Total = total
	return nil
}

// CreateAdminDictValue 创建字典值
func (ar *AuthRpc) CreateAdminDictValue(ctx context.Context, req auth.AdminDictValue, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminDict.CreateDictValue(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// UpdateAdminDictValue 更新字典值
func (ar *AuthRpc) UpdateAdminDictValue(ctx context.Context, req auth.AdminDictValue, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminDict.UpdateDictValue(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}

// DeleteAdminDictValue 删除字典值
func (ar *AuthRpc) DeleteAdminDictValue(ctx context.Context, req auth.AdminDictValue, resp *lib.Response) error {
	operator := application.GetHeader(ctx, "userName")
	if err := logic.AdminDict.DeleteDictValue(req, operator); err != nil {
		serr.Format(err, resp)
		return nil
	}
	return nil
}
