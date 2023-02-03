package auth

import "github.com/gly-hub/go-admin/common/model/lib"

type (
	LoginParams struct {
		UserId   string `json:"userId"`
		Password string `json:"password"`
	}

	LoginResp struct {
		lib.Response
		Token string `json:"token"`
	}
)

type (
	LogoutParams struct {
	}

	LogoutResp struct {
		lib.Response
	}
)

type (
	CheckTokenParams struct {
		Token string `json:"token"`
	}

	CheckTokenResp struct {
		lib.Response
		Id       int64  `json:"id"`
		UserId   string `json:"userId"`
		UserName string `json:"userName"`
		NewToken string `json:"newToken"`
	}
)
