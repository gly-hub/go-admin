package model

// AdminPermission 权限组
type AdminPermission struct {
	Id     int    `json:"id"`     // 自增键
	Title  string `json:"title"`  // 标题
	Desc   string `json:"desc"`   // 描述
	Status int    `json:"status"` // 权限组状态
	BaseModel
}

func (ap *AdminPermission) TableName() string {
	return "admin_permission"
}

// AdminPermissionMenu 权限组菜单
type AdminPermissionMenu struct {
	Id           int `json:"id"`            // 自增键
	PermissionId int `json:"permission_id"` // 权限组id
	MenuId       int `json:"menu_id"`       // 菜单id
}

func (apm *AdminPermissionMenu) TableName() string {
	return "admin_permission_menu"
}
