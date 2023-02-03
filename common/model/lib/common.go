package lib

type Response struct {
	Code int32  `json:"-"`
	Msg  string `json:"-"`
}

type PageModel struct {
	Page  int `json:"page"`  // 页码
	Limit int `json:"limit"` // 大小
}

type BaseModel struct {
	CreatedAt int64  `json:"createdAt"` // 创建时间
	UpdatedAt int64  `json:"updatedAt"` // 更新时间
	CreateBy  string `json:"createBy"`  // 创建人
	UpdateBy  string `json:"updateBy"`  // 更新人
}
