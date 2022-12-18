package model

type AdminUser struct {
	Id int32 `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (au *AdminUser) TableName() string{
	return "admin_user"
}
