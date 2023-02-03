package auth

import "github.com/gly-hub/go-admin/common/model/lib"

type AdminDictLabel struct {
	Id     int    `json:"id"`     // id
	Label  string `json:"label"`  // name
	Status int    `json:"status"` // 状态
	Desc   string `json:"desc"`   // 描述
	lib.BaseModel
	IsDelete bool `json:"isDelete"` // 是否删除
}

type AdminDictValue struct {
	Id      int    `json:"id"`      // id
	LabelId int    `json:"labelId"` // name
	Key     string `json:"type"`    // key值
	Value   string `json:"value"`   // value值
	Status  int    `json:"status"`  // 状态
	Desc    string `json:"desc"`    // 描述
	Sort    int    `json:"sort"`    // 排序
	lib.BaseModel
	IsDelete bool `json:"isDelete"` // 是否删除
}

type (
	SearchAdminDictParams struct {
		lib.PageModel
		DictLabel string `json:"dictLabel"` // 字典标签
	}

	SearchAdminDictResp struct {
		lib.Response
		Total int64            `json:"total"` // 总数
		List  []AdminDictLabel `json:"list"`  // 列表
	}
)

type (
	SearchAdminValueParams struct {
		lib.PageModel
		LabelId int `json:"labelId"` // 字典id
	}

	SearchAdminValueResp struct {
		lib.Response
		Total int64            `json:"total"` // 总数
		List  []AdminDictValue `json:"list"`  // 列表
	}
)
