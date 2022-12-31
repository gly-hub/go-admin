package logic

import (
	"github.com/gly-hub/go-admin/authorize/internal/dao"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-admin/common/model/auth"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/tools/copyx"
	"github.com/jinzhu/gorm"
	"sort"
	"time"
)

var AdminMenu adminMenuLogic

type adminMenuLogic struct {
}

// GetMenuList 获取菜单列表
func (am *adminMenuLogic) GetMenuList(in auth.SearchAdminMenuParams) ([]model.AdminMenuTree, error) {
	// 获取模糊菜单名列表
	var wheres = make(map[string]interface{})
	if in.Title != "" {
		wheres["title like ?"] = "%" + in.Title + "%"
	}

	var menuList []model.AdminMenu
	menus, err := dao.AdminMenu.GetAdminMenusByWheres(wheres)
	if err != nil {
		logger.Error(err)
		return nil, enum.DataBaseErr
	}
	menuList = append(menuList, menus...)

	for _, menu := range menus {
		// 向上获取菜单列表
		menuList = append(menuList, am.getAdminMenuUp(menu.ParentId)...)
		// 向下获取菜单列表
		menuList = append(menuList, am.getAdminMenuDown(menu.Id)...)
	}

	// 菜单去重
	var (
		newMenuList []model.AdminMenu
		menuMap     = make(map[int]bool)
	)
	for _, menu := range menuList {
		if _, ok := menuMap[menu.Id]; ok {
			continue
		}
		newMenuList = append(newMenuList, menu)
		menuMap[menu.Id] = true
	}

	// 构建树
	data := am.recursiveSetMenu(newMenuList, 0)

	return data, nil
}

// getAdminMenuUp 向上获取父级菜单
func (am *adminMenuLogic) getAdminMenuUp(parentId int) (data []model.AdminMenu) {
	menu, err := dao.AdminMenu.GetAdminMenuById(parentId)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return
	}

	if err == gorm.ErrRecordNotFound {
		return
	}

	data = append(data, menu)
	data = append(data, am.getAdminMenuUp(menu.ParentId)...)
	return
}

// getAdminMenuDown 向下获取子级菜单
func (am *adminMenuLogic) getAdminMenuDown(parentId int) (data []model.AdminMenu) {
	menus, err := dao.AdminMenu.GetAdminMenusByParentId(parentId)
	if err != nil {
		logger.Error(err)
		return
	}
	data = append(data, menus...)
	for _, menu := range menus {
		sonMenus := am.getAdminMenuDown(menu.Id)
		data = append(data, sonMenus...)
	}

	return
}

// CreateMenu 创建菜单
func (am *adminMenuLogic) CreateMenu(in auth.AdminMenu, operator string) error {
	// 首先判断菜单类型
	if in.ParentId == 0 && in.MenuType != enum.MenuTypeModule {
		return enum.MenuNeedModule
	}

	// 校验父级是否存在
	if in.ParentId != 0 {
		_, err := dao.AdminMenu.GetAdminMenuById(in.ParentId)
		if err != nil && err != gorm.ErrRecordNotFound {
			logger.Error(err)
			return enum.DataBaseErr
		}
		if err == gorm.ErrRecordNotFound {
			return enum.MenuParentMenuNotFound
		}
	}

	// 插入数据
	var menu model.AdminMenu
	if err := copyx.DeepCopy(&menu, &in); err != nil {
		logger.Error(err)
		return enum.SystemErr
	}
	menu.Id = 0
	menu.CreateBy = operator
	menu.CreatedAt = time.Now().Unix()
	if err := dao.AdminMenu.InsertAdminMenu(menu); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}

	return nil
}

// UpdateMenu 更新菜单
func (am *adminMenuLogic) UpdateMenu(in auth.AdminMenu, operator string) error {
	// 先判断主键
	if in.Id == 0 {
		return enum.PrimaryKeyNotFound
	}

	// 校验菜单是否存在
	_, err := dao.AdminMenu.GetAdminMenuById(in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	if err == gorm.ErrRecordNotFound {
		logger.Error("菜单不存在")
		return enum.PrimaryKeyNotFound
	}

	// 首先判断菜单类型
	if in.ParentId == 0 && in.MenuType != enum.MenuTypeModule {
		return enum.MenuNeedModule
	}

	// 校验父级是否存在
	if in.ParentId != 0 {
		_, err := dao.AdminMenu.GetAdminMenuById(in.ParentId)
		if err != nil && err != gorm.ErrRecordNotFound {
			logger.Error(err)
			return enum.DataBaseErr
		}
		if err == gorm.ErrRecordNotFound {
			return enum.MenuParentMenuNotFound
		}
	}

	// 插入数据
	var menu model.AdminMenu
	if err := copyx.DeepCopy(&menu, &in); err != nil {
		logger.Error(err)
		return enum.SystemErr
	}
	menu.UpdateBy = operator
	menu.UpdatedAt = time.Now().Unix()
	if err := dao.AdminMenu.UpdateAdminMenu(menu); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}

	return nil
}

// DeleteMenu 删除菜单
func (am *adminMenuLogic) DeleteMenu(in auth.AdminMenu, operator string) error {
	// 校验菜单是否存在
	_, err := dao.AdminMenu.GetAdminMenuById(in.Id)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error(err)
		return enum.DataBaseErr
	}

	if err == gorm.ErrRecordNotFound {
		return enum.MenuNotFound
	}

	var menu model.AdminMenu
	if err := copyx.DeepCopy(&menu, &in); err != nil {
		logger.Error(err)
		return enum.SystemErr
	}
	menu.IsDelete = true
	menu.UpdateBy = operator
	if err := dao.AdminMenu.UpdateAdminMenu(menu); err != nil {
		logger.Error(err)
		return enum.DataBaseErr
	}

	return nil
}

// GetMenuTreeByUserId 通过userId获取菜单树
func (am *adminMenuLogic) GetMenuTreeByUserId(userId string, platform int) (data []model.AdminMenuTree, err error) {
	data, err = am.getMenuByUserId(userId, platform)
	sort.Sort(model.AdminMenuSlice(data))
	return
}

// getMenuByUserId 获取用户菜单列表
func (am *adminMenuLogic) getMenuByUserId(userId string, platform int) (data []model.AdminMenuTree, err error) {
	menus := make([]model.AdminMenu, 0)

	// 校验是否超级权限
	if dao.AdminUser.IsAdmin(userId) {
		// 获取所有的菜单列表
		menus, err = dao.AdminMenu.GetAdminMenuByPlatform(platform)
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
		menus, err = dao.AdminPermission.GetAdminMenuByPermissionIds(permissionIds, platform)
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
