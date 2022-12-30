package dao

import (
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/logger"
)

var AdminUser adminUserDao

type adminUserDao struct {
	application.DB
}

// GetAdminUserInfo 获取用户信息
func (a adminUserDao) GetAdminUserInfo(userId string) (userInfo model.AdminUser, err error) {
	err = a.GetRDB().Model(model.AdminUser{}).Where("user_id = ?", userId).Find(&userInfo).Error
	return
}

// IsAdmin 校验用户是否存在超级权限
func (a adminUserDao) IsAdmin(userId string) bool {
	var count int64
	err := a.GetRDB().Model(model.AdminUserPermission{}).Joins(
		"left join admin_permission on admin_permission.id = admin_user_permission.permission_id").
		Where("admin_user_permission.user_id = ?", userId).
		Where("admin_permission.title = ?", "admin").Count(&count).Error
	if err != nil {
		logger.Error(err)
		return false
	}

	if count > 0 {
		return true
	}

	return false
}

// GetUserPermissionIds 获取用户权限组
func (a adminUserDao) GetUserPermissionIds(userId string) ([]int, error) {
	var permissionIds []int
	err := a.GetRDB().Model(model.AdminUserPermission{}).Where("user_id = ?", userId).Pluck("permission_id", permissionIds).Error
	return permissionIds, err
}
