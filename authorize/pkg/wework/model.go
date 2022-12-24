package wework

type Error struct {
	ErrCode int    `json:"errorcode"` // 出错返回码，为0表示成功，非0表示调用失败
	ErrMsg  string `json:"errmsg"`    // 返回码提示语
}

type getAccessTokenResp struct {
	AccessToken string `json:"access_token"` // 获取到的凭证，最长为512字节
	ExpiresIn   int32  `json:"expires_in"`   // 凭证的有效时间（秒）
}

type UserSimple struct {
	UserId     string `json:"userid"`      // 成员UserID。对应管理端的帐号
	Name       string `json:"name"`        // 成员名称
	Department []int  `json:"department"`  // 成员所属部门列表
	OpenUserId string `json:"open_userid"` // 全局唯一
}

type getUserSimpleListResp struct {
	UserList []UserSimple `json:"userlist"` // 成员列表
}

type User struct {
	UserId     string `json:"userid"`      // 成员UserID
	Name       string `json:"name"`        // 成员名称
	Alias      string `json:"alias"`       // 别名
	Mobile     string `json:"mobile"`      // 手机号码
	Department []int  `json:"department"`  // 成员所属部门id列表
	Order      []int  `json:"order"`       // 部门内的排序值，默认为0
	Gender     string `json:"gender"`      // 性别。0表示未定义，1表示男性，2表示女性。
	Email      string `json:"email"`       // 邮箱
	BizMail    string `json:"biz_mail"`    // 企业邮箱
	Avatar     string `json:"avatar"`      // 头像url
	OpenUserId string `json:"open_userid"` // 全局唯一
	Status     int    `json:"status"`      // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业。
}

type getUserListResp struct {
	UserList []User `json:"userlist"` // 成员列表
}

type Department struct {
	Id               int      `json:"id"`                // 部门id
	Name             string   `json:"name"`              // 部门名称
	NameEn           string   `json:"name_en"`           // 部门英文名称
	DepartmentLeader []string `json:"department_leader"` // 部门负责人的UserID
	ParentId         int      `json:"parentid"`          // 父部门id
	Order            int      `json:"order"`             // 在父部门中的次序值
}

type getDepartmentListResp struct {
	Department []Department `json:"department"` // 成员列表
}

type DepartmentSimple struct {
	Id       int `json:"id"`       // 部门id
	ParentId int `json:"parentid"` // 父部门id
	Order    int `json:"order"`    // 在父部门中的次序值
}

type getDepartmentSimpleListResp struct {
	Department []DepartmentSimple `json:"department_id"` // 成员列表
}
