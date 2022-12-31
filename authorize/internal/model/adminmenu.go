package model

type AdminMenuTree struct {
	AdminMenu
	ChildMenus []AdminMenuTree
}

type AdminMenu struct {
	Id         int    `json:"id"`         // 自增键
	MenuName   string `json:"name"`       // 菜单名
	Title      string `json:"title"`      // 标题
	Icon       string `json:"icon"`       // 图标
	Path       string `json:"path"`       // 路径
	Paths      string `json:"paths"`      // 路径树
	MenuType   string `json:"menu_type"`  // 菜单类型 P:模块 M：菜单 C：页面 T：tab F：按钮
	Action     string `json:"action"`     // 行为
	Permission string `json:"permission"` // 权限
	ParentId   int    `json:"parent_id"`  // 父级id
	NoCache    bool   `json:"no_cache"`   // 是否缓存
	Component  string `json:"component"`  // 组件
	Sort       int    `json:"sort"`       // 排序
	Visible    bool   `json:"visible"`    // 是否可见
	IsFrame    bool   `json:"is_frame"`   // 是否是frame
	BaseModel
	IsDelete bool `json:"is_delete"` // 是否删除
}

func (am *AdminMenu) TableName() string {
	return "admin_menus"
}

type AdminMenuSlice []AdminMenuTree

func (x AdminMenuSlice) Len() int           { return len(x) }
func (x AdminMenuSlice) Less(i, j int) bool { return x[i].Sort < x[j].Sort }
func (x AdminMenuSlice) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
