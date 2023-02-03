package auth

import (
	"github.com/gly-hub/go-admin/common/model/lib"
)

type UserMenuTreeParams struct {
}

type AdminMenuTree struct {
	AdminMenu
	ChildMenus []AdminMenuTree `json:"children"`
}

type AdminMenu struct {
	Id         int    `json:"id"`         // 自增键
	MenuName   string `json:"name"`       // 菜单名
	Title      string `json:"title"`      // 标题
	Icon       string `json:"icon"`       // 图标
	Platform   int    `json:"platform"`   // 所属模块
	Path       string `json:"path"`       // 路径
	Paths      string `json:"paths"`      // 路径树
	MenuType   string `json:"menuType"`   // 菜单类型 P:模块 M：菜单 C：页面 T：tab F：按钮
	Action     string `json:"action"`     // 行为
	Permission string `json:"permission"` // 权限
	ParentId   int    `json:"parentId"`   // 父级id
	NoCache    bool   `json:"noCache"`    // 是否缓存
	Component  string `json:"component"`  // 组件
	Sort       int    `json:"sort"`       // 排序
	Visible    bool   `json:"visible"`    // 是否可见
	IsFrame    bool   `json:"isFrame"`    // 是否是frame
	CreateBy   string `json:"createBy"`   // 创建人
	UpdateBy   string `json:"updateBy"`   // 更新人
	CreatedAt  int64  `json:"createdAt"`  // 创建时间
	UpdatedAt  int64  `json:"updatedAt"`  // 更新时间
	IsDelete   bool   `json:"isDelete"`   // 是否删除
}

type UserMenuTreeResp struct {
	lib.Response
	Menus []AdminMenuTree `json:"menus"`
}

type SearchAdminMenuParams struct {
	Title string `json:"title"` // 菜单名。模糊查询
}
