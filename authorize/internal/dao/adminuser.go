package dao

import (
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/application"
)

var AdminUser adminUserDao

type adminUserDao struct {
	application.DB
}

func (a adminUserDao) GetAdminUserInfo(userName string) (userInfo model.AdminUser, err error) {
	err = a.GetRDB().Model(model.AdminUser{}).Where("user_name = ?", userName).Find(&userInfo).Error
	return
}
