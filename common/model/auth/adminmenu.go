package auth

import "github.com/gly-hub/go-admin/common/model/common"

type UserMenuTreeParams struct {
}

type AdminMenuTree struct {
	AdminMenu
	ChildMenus []AdminMenuTree
}

type AdminMenu struct {
	Id         int    `json:"id"`         // 自增键
	MenuName   string `json:"name"`       // 菜单名
	Title      string `json:"title"`      // 标题
	Icon       string `json:"icon"`       // 图标
	Platform   int    `json:"platform"`   // 所属模块
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
	CreateBy   string `json:"create_by"`  // 创建人
	UpdateBy   string `json:"update_by"`  // 更新人
	CreatedAt  int64  `json:"created_at"` // 创建时间
	UpdatedAt  int64  `json:"updated_at"` // 更新时间
	DeletedAt  int64  `json:"deleted_at"` // 删除时间
	IsDelete   bool   `json:"is_delete"`  // 是否删除
}

type UserMenuTreeResp struct {
	common.Response
	Menus []AdminMenuTree `json:"menus"`
}
