package model

type AdminUser struct {
	Id         int32  `json:"id"`
	UserId     string `json:"userid"`      // 成员UserID
	Password   string `json:"password"`    // 密码
	Name       string `json:"name"`        // 成员名称
	Alias      string `json:"alias"`       // 别名
	Mobile     string `json:"mobile"`      // 手机号码
	Gender     string `json:"gender"`      // 性别。0表示未定义，1表示男性，2表示女性。
	Email      string `json:"email"`       // 邮箱
	BizMail    string `json:"biz_mail"`    // 企业邮箱
	Avatar     string `json:"avatar"`      // 头像url
	OpenUserId string `json:"open_userid"` // 全局唯一
	Status     int    `json:"status"`      // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。
	Source     int    `json:"source"`      // 来源：1=平台创建 2=企微 3=钉钉
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间
	CreateBy   string `json:"create_by"`   // 创建人
	UpdateBy   string `json:"update_by"`   // 更新人
}

func (au *AdminUser) TableName() string {
	return "admin_user"
}

type AdminUserDepartment struct {
	Id           int `json:"id"`            // 自增键
	UserId       int `json:"user_id"`       // 用户id
	DepartmentId int `json:"department_id"` // 部门id
	Order        int `json:"order"`         // 排序
}

func (au *AdminUserDepartment) TableName() string {
	return "admin_user_department"
}

type AdminUserMenu struct {
	Id     int `json:"id"`      // 自增键
	UserId int `json:"user_id"` // 用户id
	MenuId int `json:"menu_id"` // 菜单id
}

func (aup *AdminUserMenu) TableName() string {
	return "admin_user_permission"
}
