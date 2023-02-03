package global

import (
	"github.com/gly-hub/go-admin/authorize/tools/jwt"
)

var JwtGlobal *jwt.Jwt

func InitJwt() {
	JwtGlobal = jwt.NewJwt(BaseConfig.Jwt.Key, BaseConfig.Jwt.Expire)
}
