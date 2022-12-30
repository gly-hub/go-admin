package model

type BaseModel struct {
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
	CreateBy  string `json:"create_by"`  // 创建人
	UpdateBy  string `json:"update_by"`  // 更新人
}
