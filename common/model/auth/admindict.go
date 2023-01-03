package auth

import "github.com/gly-hub/go-admin/common/model/common"

type AdminDictLabel struct {
	Id     int    `json:"id"`     // id
	Label  string `json:"label"`  // name
	Status int    `json:"status"` // 状态
	Desc   string `json:"desc"`   // 描述
	common.BaseModel
	IsDelete bool `json:"is_delete"` // 是否删除
}

type AdminDictValue struct {
	Id      int    `json:"id"`       // id
	LabelId int    `json:"label_id"` // name
	Key     string `json:"type"`     // key值
	Value   string `json:"value"`    // value值
	Status  int    `json:"status"`   // 状态
	Desc    string `json:"desc"`     // 描述
	Sort    int    `json:"sort"`     // 排序
	common.BaseModel
	IsDelete bool `json:"is_delete"` // 是否删除
}

type (
	SearchAdminDictParams struct {
		common.PageModel
		DictLabel string `json:"dict_label"` // 字典标签
	}

	SearchAdminDictResp struct {
		common.Response
		Total int64            `json:"total"` // 总数
		List  []AdminDictLabel `json:"list"`  // 列表
	}
)

type (
	SearchAdminValueParams struct {
		common.PageModel
		LabelId int `json:"label_id"` // 字典id
	}

	SearchAdminValueResp struct {
		common.Response
		Total int64            `json:"total"` // 总数
		List  []AdminDictValue `json:"list"`  // 列表
	}
)
