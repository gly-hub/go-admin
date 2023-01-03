package model

type AdminDictLabel struct {
	Id     int    `json:"id"`     // id
	Label  string `json:"label"`  // name
	Status int    `json:"status"` // 状态
	Desc   string `json:"desc"`   // 描述
	BaseModel
	IsDelete bool `json:"is_delete"` // 是否删除
}

func (adl *AdminDictLabel) TableName() string {
	return "admin_dict_label"
}

type AdminDictValue struct {
	Id      int    `json:"id"`       // id
	LabelId int    `json:"label_id"` // name
	Key     string `json:"type"`     // key值
	Value   string `json:"value"`    // value值
	Status  int    `json:"status"`   // 状态
	Desc    string `json:"desc"`     // 描述
	Sort    int    `json:"sort"`     // 排序
	BaseModel
	IsDelete bool `json:"is_delete"` // 是否删除
}

func (adv *AdminDictValue) TableName() string {
	return "admin_dict_value"
}
