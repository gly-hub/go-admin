package dao

import (
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/application"
)

var AdminMenu adminMenuDao

type adminMenuDao struct {
	application.DB
}

// GetAllAdminMenu 获取所有菜单
func (amd adminMenuDao) GetAllAdminMenu() (menus []model.AdminMenu, err error) {
	err = amd.GetRDB().Model(model.AdminMenu{}).Where("deleted_at is null").Find(&menus).Error
	return
}
