package logic

import (
	"github.com/gly-hub/go-admin/authorize/internal/dao"
	"github.com/gly-hub/go-admin/authorize/internal/model"
)

var AdminDepartment adminDepartmentLogic

type adminDepartmentLogic struct {
}

//获取部门结构树
func (logic *adminDepartmentLogic) GetDepartmentTree() (tree []model.AdminDepartment, err error) {
	var list []model.AdminDepartment
	list, err = dao.AdminDepartment.GetDepartmentList()
	if err != nil {
		return
	}
	//递归回去结构树
	tree = logic.getDepartmentTree(list, 0)

	return
}

//生成结构树
func (logic *adminDepartmentLogic) getDepartmentTree(list []model.AdminDepartment, pid int) (tree []model.AdminDepartment) {
	// 递归结构树
	for _, v := range list {
		if v.ParentId == pid {
			v.Children = logic.getDepartmentTree(list, v.Id)
			tree = append(tree, v)
		}
	}
	return
}
