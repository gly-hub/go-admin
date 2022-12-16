package auth

import (
	"github.com/gly-hub/go-admin/internal-gateway/internal/service/auth"
	routing "github.com/qiangxue/fasthttp-routing"
)

func InitAuthRoute(baseRouter *routing.RouteGroup){
	authHandler := auth.AuthController{}

	// 登录登出
	baseRouter.Post("/login", authHandler.Login)

	authGroup := baseRouter.Group("/auth")
	{
		// 系统路由
		systemGroup := authGroup.Group("/system")
		{
			// 获取菜单树
			systemGroup.Post("/menu/tree", nil)
		}
	}
}
