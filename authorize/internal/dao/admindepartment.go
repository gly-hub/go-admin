package dao

import (
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/application"
)

var AdminDepartment adminDepartmentDao

type adminDepartmentDao struct {
	application.DB
	application.Redis
}

//获取部门列表
func (dao *adminDepartmentDao) GetDepartmentList() (list []model.AdminDepartment, err error) {
	err = dao.GetDB().Model(model.AdminDepartment{}).Find(&list).Error
	return
}
