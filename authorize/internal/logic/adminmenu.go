package logic

type AdminMenu adminMenuLogic

type adminMenuLogic struct {
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
func (am *adminMenuLogic) GetMenuTreeByUserId(userId string) {

}
