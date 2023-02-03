package rpcserver

const (
	// ServerName 服务名
	ServerName = "authorize"
)

// 方法名
const (
	Login      = "Login"      // 登录
	Logout     = "Logout"     // 注销
	CheckToken = "CheckToken" // 校验token
	UserMenu   = "UserMenu"   // 获取用户菜单树
	SystemInfo = "SystemInfo" // 获取系统信息
)

// 菜单管理
const (
	SearchAdminMenu = "SearchAdminMenu" // 查询菜单列表
	CreateAdminMenu = "CreateAdminMenu" // 创建菜单项
	UpdateAdminMenu = "UpdateAdminMenu" // 更新菜单项
	DeleteAdminMenu = "DeleteAdminMenu" // 删除菜单项
)

// 字典管理
const (
	SearchAdminDictLabel = "SearchAdminDictLabel" // 查询字典列表
	CreateAdminDict      = "CreateAdminDict"      // 创建字典标签
	UpdateAdminDict      = "UpdateAdminDict"      // 更新字典标签
	DeleteAdminDict      = "DeleteAdminDict"      // 删除字典标签
	SearchAdminDictValue = "SearchAdminDictValue" // 查询字典值列表
	CreateAdminDictValue = "CreateAdminDictValue" // 创建字典值
	UpdateAdminDictValue = "UpdateAdminDictValue" // 更新字典值
	DeleteAdminDictValue = "DeleteAdminDictValue" // 删除字典值
)
