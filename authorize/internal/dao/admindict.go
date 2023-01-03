package dao

import (
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/application"
)

var AdminDict adminDictDao

type adminDictDao struct {
	application.DB
}

// GetAdminDictLabels 获取字典标签列表
func (ad adminDictDao) GetAdminDictLabels(dictLabel string, page, limit int) (labels []model.AdminDictLabel, total int64, err error) {
	tx := ad.GetRDB().Model(model.AdminDictLabel{}).Where("is_delete = false")
	if dictLabel != "" {
		tx = tx.Where("label like ?", "%"+dictLabel+"%")
	}
	err = tx.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&labels).Error
	return
}

// GetAdminDictLabelById 通过id获取字典标签
func (ad adminDictDao) GetAdminDictLabelById(id int) (label model.AdminDictLabel, err error) {
	err = ad.GetRDB().Model(model.AdminDictLabel{}).Where("id = ?", id).First(&label).Error
	return
}

// GetAdminDictLabelByLabel 通过label获取字典标签
func (ad adminDictDao) GetAdminDictLabelByLabel(labelStr string) (label model.AdminDictLabel, err error) {
	err = ad.GetRDB().Model(model.AdminDictLabel{}).Where("label = ?", labelStr).First(&label).Error
	return
}

// CreateAdminDictLabel 创建字典标签
func (ad adminDictDao) CreateAdminDictLabel(label model.AdminDictLabel) (err error) {
	err = ad.GetWDB().Model(model.AdminDictLabel{}).Create(&label).Error
	return
}

// UpdateAdminDictLabel 更新字典标签
func (ad adminDictDao) UpdateAdminDictLabel(label model.AdminDictLabel) (err error) {
	err = ad.GetWDB().Model(model.AdminDictLabel{}).Where("id = ?", label.Id).Save(&label).Error
	return
}

// GetAdminDictValues 获取字典值列表
func (ad adminDictDao) GetAdminDictValues(dictId, page, limit int) (labels []model.AdminDictValue, total int64, err error) {
	tx := ad.GetRDB().Model(model.AdminDictValue{}).Where("label_id = ?", dictId).Where("is_delete = false")
	err = tx.Count(&total).Limit(limit).Offset((page - 1) * limit).Find(&labels).Error
	return
}

// CreateAdminDictValue 创建字典标签
func (ad adminDictDao) CreateAdminDictValue(label model.AdminDictValue) (err error) {
	err = ad.GetWDB().Model(model.AdminDictValue{}).Create(&label).Error
	return
}

// UpdateAdminDictValue 更新字典标签
func (ad adminDictDao) UpdateAdminDictValue(label model.AdminDictValue) (err error) {
	err = ad.GetWDB().Model(model.AdminDictValue{}).Where("id = ?", label.Id).Save(&label).Error
	return
}

// CheckAdminDictValue 校验是否存在相同字典值
func (ad adminDictDao) CheckAdminDictValue(labelId int, key, value string) (dictValue model.AdminDictValue, err error) {
	err = ad.GetRDB().Model(model.AdminDictValue{}).Where("label_id = ?", labelId).
		Where("key = ? or value = ?", key, value).First(&dictValue).Error
	return
}

// GetAdminDictValueById 通过id获取字典值
func (ad adminDictDao) GetAdminDictValueById(valueId int) (dictValue model.AdminDictValue, err error) {
	err = ad.GetRDB().Model(model.AdminDictValue{}).Where("id = ?", valueId).
		First(&dictValue).Error
	return
}
