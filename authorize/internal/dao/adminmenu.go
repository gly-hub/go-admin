package dao

import "github.com/gly-hub/go-dandelion/application"

var AdminMenu adminMenuDao

type adminMenuDao struct {
	application.DB
}

func (amd adminMenuDao) GetAdminMenu() {
	//amd.GetRDB().
}
