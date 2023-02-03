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
	err = amd.GetRDB().Model(model.AdminMenu{}).Where("is_delete = 0").Find(&menus).Error
	return
}

// GetAdminMenuByPlatform 获取平台菜单
func (amd adminMenuDao) GetAdminMenuByPlatform(platform int) (menus []model.AdminMenu, err error) {
	err = amd.GetRDB().Model(model.AdminMenu{}).Where("is_delete = 0 and platform = ? and menu_type in ('M', 'C')", platform).Find(&menus).Error
	return
}

// GetAdminMenuById 获取菜单详情
func (amd adminMenuDao) GetAdminMenuById(menuId int) (menu model.AdminMenu, err error) {
	err = amd.GetRDB().Model(model.AdminMenu{}).Where("is_delete = 0").Where("id = ?", menuId).First(&menu).Error
	return
}

// GetAdminMenusByParentId 通过父级id获取菜单列表
func (amd adminMenuDao) GetAdminMenusByParentId(parentId int) (menus []model.AdminMenu, err error) {
	err = amd.GetRDB().Model(model.AdminMenu{}).Where("is_delete = 0").Where("parent_id = ?", parentId).Find(&menus).Error
	return
}

// GetAdminMenusByWheres 条件获取菜单列表
func (amd adminMenuDao) GetAdminMenusByWheres(wheres map[string]interface{}) (menus []model.AdminMenu, err error) {
	tx := amd.GetRDB().Model(model.AdminMenu{})
	for k, v := range wheres {
		tx = tx.Where(k, v)
	}
	err = tx.Where("is_delete = 0").Find(&menus).Error
	return
}

// InsertAdminMenu 插入菜单
func (amd adminMenuDao) InsertAdminMenu(menu model.AdminMenu) error {
	return amd.GetWDB().Model(model.AdminMenu{}).Create(&menu).Error
}

// UpdateAdminMenu 更新菜单
func (amd adminMenuDao) UpdateAdminMenu(menu model.AdminMenu) error {
	return amd.GetWDB().Model(model.AdminMenu{}).Where("id = ?", menu.Id).Save(&menu).Error
}
