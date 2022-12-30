package logic

import (
	"github.com/gly-hub/go-admin/authorize/internal/dao"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/tools/copyx"
	"sort"
)

var AdminMenu adminMenuLogic

type adminMenuLogic struct {
}

// GetPlatformList 获取模块列表
func (am *adminMenuLogic) GetPlatformList() {

}

// GetMenuList 获取菜单列表
func (am *adminMenuLogic) GetMenuList() {

}

// CreateMenu 创建菜单
func (am *adminMenuLogic) CreateMenu() {
	// 查看父级菜单是否拥有权限 且父级是否是菜单类型

}

// UpdateMenu 更新菜单
func (am *adminMenuLogic) UpdateMenu() {
	// 查看父级菜单是否拥有权限 且父级是否是菜单类型

}

// DeleteMenu 删除菜单
func (am *adminMenuLogic) DeleteMenu() {
	// 查看父级菜单是否拥有权限

}

// GetMenuTreeByUserId 通过userId获取菜单树
func (am *adminMenuLogic) GetMenuTreeByUserId(userId string, platform int) (data []model.AdminMenuTree, err error) {
	data, err = am.getMenuByUserId(userId)
	sort.Sort(model.AdminMenuSlice(data))
	return
}

// getMenuByUserId 获取用户菜单列表
func (am *adminMenuLogic) getMenuByUserId(userId string) (data []model.AdminMenuTree, err error) {
	menus := make([]model.AdminMenu, 0)

	// 校验是否超级权限
	if dao.AdminUser.IsAdmin(userId) {
		// 获取所有的菜单列表
		menus, err = dao.AdminMenu.GetAllAdminMenu()
		if err != nil {
			logger.Error(err)
			err = enum.DataBaseErr
			return
		}
	} else {
		// 获取用户权限组
		var permissionIds []int
		permissionIds, err = dao.AdminUser.GetUserPermissionIds(userId)
		if err != nil {
			logger.Error(err)
			err = enum.DataBaseErr
			return
		}
		// 通过用户所对应才对获取
		menus, err = dao.AdminPermission.GetAdminMenuByPermissionIds(permissionIds)
		if err != nil {
			logger.Error(err)
			err = enum.DataBaseErr
			return
		}
	}

	data = am.recursiveSetMenu(menus, 0)
	return
}

// recursiveSetMenu 构建菜单树
func (am *adminMenuLogic) recursiveSetMenu(menus []model.AdminMenu, parentId int) (childMenus []model.AdminMenuTree) {
	for i, _ := range menus {
		if menus[i].ParentId == parentId {
			var child model.AdminMenuTree
			err := copyx.DeepCopy(&child, &menus[i])
			if err != nil {
				logger.Error(err)
				continue
			}
			child.ChildMenus = am.recursiveSetMenu(menus, child.Id)
			childMenus = append(childMenus, child)
		}
	}
	return childMenus
}
