package model

type AdminDepartment struct {
	Id       int    `json:"id"`        // 自增键
	Name     string `json:"name"`      // 部门名称
	NameEn   string `json:"name_en"`   // 部门英文名称
	ParentId int    `json:"parent_id"` // 父级部门id
	Order    int    `json:"order"`     // 排序
}

func (ad *AdminDepartment) TableName() string {
	return "admin_department"
}

type AdminDepartmentLeader struct {
	Id           int `json:"id"`            // 自增键
	DepartmentId int `json:"department_id"` // 部门id
	UserId       int `json:"user_id"`       // 用户id
}

func (adl *AdminDepartmentLeader) TableName() string {
	return "admin_department_leader"
}
