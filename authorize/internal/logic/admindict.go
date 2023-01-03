package logic

import (
	"github.com/gly-hub/go-admin/authorize/internal/dao"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/jinzhu/gorm"
	"time"
)

var AdminDict adminDictLogic

type adminDictLogic struct {
}

// GetDictList 获取字典标签列表
func (adl adminDictLogic) GetDictList(in auth.SearchAdminDictParams) ([]model.AdminDictLabel, int64, error) {
	list, total, err := dao.AdminDict.GetAdminDictLabels(in.DictLabel, in.Page, in.Limit)
	if err != nil {
		logger.Error(err)
		return nil, 0, enum.DataBaseErr
	}
	return list, total, nil
}

// CreateDict 创建字典
func (adl adminDictLogic) CreateDict(in auth.AdminDictLabel, operator string) error {
	var label = model.AdminDictLabel{
		Label:  in.Label,
		Status: in.Status,
		Desc:   in.Desc,
		BaseModel: model.BaseModel{
			CreatedAt: time.Now().Unix(),
			CreateBy:  operator,
		},
		IsDelete: false,
	}
	err := dao.AdminDict.CreateAdminDictLabel(label)
	if err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}
	return nil
}

// UpdateDict 更新字典
func (adl adminDictLogic) UpdateDict(in auth.AdminDictLabel, operator string) error {
	// 获取标签
	label, err := dao.AdminDict.GetAdminDictLabelById(in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	if err == gorm.ErrRecordNotFound {
		return enum.PrimaryKeyNotFound
	}

	// 校验标签是否重复
	if labelOld, err := dao.AdminDict.GetAdminDictLabelByLabel(in.Label); err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	} else if err != gorm.ErrRecordNotFound && labelOld.Id != in.Id {
		return enum.DictLabelRepeat
	}

	label.Label = in.Label
	label.Status = in.Status
	label.Desc = in.Desc
	label.UpdateBy = operator
	label.UpdatedAt = time.Now().Unix()

	if err := dao.AdminDict.CreateAdminDictLabel(label); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}
	return nil
}

// DeleteDict 删除字典
func (adl adminDictLogic) DeleteDict(in auth.AdminDictLabel, operator string) error {
	// 获取标签
	label, err := dao.AdminDict.GetAdminDictLabelById(in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	if err == gorm.ErrRecordNotFound {
		return enum.PrimaryKeyNotFound
	}

	label.UpdateBy = operator
	label.UpdatedAt = time.Now().Unix()
	label.IsDelete = true

	if err := dao.AdminDict.CreateAdminDictLabel(label); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}
	return nil
}

// GetDictValueList 获取字典值列表
func (adl adminDictLogic) GetDictValueList(in auth.SearchAdminValueParams) ([]model.AdminDictValue, int64, error) {
	list, total, err := dao.AdminDict.GetAdminDictValues(in.LabelId, in.Page, in.Limit)
	if err != nil {
		logger.Error(err)
		return nil, 0, enum.DataBaseErr
	}
	return list, total, nil
}

// CreateDictValue 创建字典值
func (adl adminDictLogic) CreateDictValue(in auth.AdminDictValue, operator string) error {
	// 校验父级是否存在
	// 获取标签
	_, err := dao.AdminDict.GetAdminDictLabelById(in.LabelId)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	if err == gorm.ErrRecordNotFound {
		return enum.PrimaryKeyNotFound
	}

	// 校验同一个标签下key或者value是否存在
	if _, err := dao.AdminDict.CheckAdminDictValue(in.LabelId, in.Key, in.Value); err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	var value = model.AdminDictValue{
		LabelId: in.LabelId,
		Key:     in.Key,
		Value:   in.Value,
		Status:  in.Status,
		Desc:    in.Desc,
		Sort:    in.Sort,
		BaseModel: model.BaseModel{
			CreatedAt: time.Now().Unix(),
			CreateBy:  operator,
		},
		IsDelete: false,
	}
	if err := dao.AdminDict.CreateAdminDictValue(value); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}
	return nil
}

// UpdateDictValue 更新字典值
func (adl adminDictLogic) UpdateDictValue(in auth.AdminDictValue, operator string) error {
	if in.Key == "" || in.Value == "" {
		return enum.DictValueParamsErr
	}

	// 校验标签值是否存在
	value, vErr := dao.AdminDict.GetAdminDictValueById(in.Id)
	if vErr != nil && vErr != gorm.ErrRecordNotFound {
		return enum.DataBaseErr
	} else if vErr == gorm.ErrRecordNotFound {
		return enum.PrimaryKeyNotFound
	}

	// 校验标签是否重复
	// 校验同一个标签下key或者value是否存在
	if valueOld, err := dao.AdminDict.CheckAdminDictValue(value.LabelId, in.Key, in.Value); err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	} else if err != gorm.ErrRecordNotFound && valueOld.Id != in.Id {
		return enum.DictValueIsExist
	}

	value.Key = in.Key
	value.Value = in.Value
	value.Status = in.Status
	value.Desc = in.Desc
	value.Sort = in.Sort
	value.UpdateBy = operator
	value.UpdatedAt = time.Now().Unix()

	if err := dao.AdminDict.CreateAdminDictValue(value); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}
	return nil
}

// DeleteDictValue 删除字典值
func (adl adminDictLogic) DeleteDictValue(in auth.AdminDictValue, operator string) error {
	// 获取字典值
	value, err := dao.AdminDict.GetAdminDictValueById(in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	if err == gorm.ErrRecordNotFound {
		return enum.PrimaryKeyNotFound
	}

	value.UpdateBy = operator
	value.UpdatedAt = time.Now().Unix()
	value.IsDelete = true

	if err := dao.AdminDict.CreateAdminDictValue(value); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}
	return nil
}
