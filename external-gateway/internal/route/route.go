package route

import "github.com/gly-hub/go-admin/external-gateway/internal/route/auth"

func InitRoute(){
	// auth服务相关路由
	auth.InitAuthRoute()
}
