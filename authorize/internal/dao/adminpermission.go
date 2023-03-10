package dao

import (
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/application"
)

var AdminPermission adminPermissionDao

type adminPermissionDao struct {
	application.DB
}

// GetAdminMenuByPermissionIds 通过用户id获取菜单列表
func (amd adminPermissionDao) GetAdminMenuByPermissionIds(permissionIds []int, platform int) (menus []model.AdminMenu, err error) {
	err = amd.GetDB().Model(model.AdminMenu{}).Joins(
		"left join admin_permission_menu on admin_permission_menu.menu_id = admin_menus.id").
		Where("admin_menus.deleted_at is null and admin_menus.platform = ?", platform).
		Where("admin_permission_menu.permission_id in (?)", permissionIds).Find(&menus).Error
	return
}
