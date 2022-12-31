package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/internal-gateway/internal/service/auth"
)

func InitAuthRoute(baseRouter *routing.RouteGroup) {
	authHandler := auth.AuthController{}

	// 登录登出
	baseRouter.Post("/login", authHandler.Login)

	authGroup := baseRouter.Group("/auth")
	{
		// 系统路由
		systemGroup := authGroup.Group("/system")
		{
			// 获取菜单树
			systemGroup.Get("/menu/tree", authHandler.GetMenuTree)
		}

		// 菜单管理
		menuGroup := authGroup.Group("/menu")
		{
			menuGroup.Post("/search", authHandler.GetMenuList)
			menuGroup.Post("", authHandler.CreateMenu)
			menuGroup.Put("", authHandler.UpdateMenu)
			menuGroup.Delete("/<id>", authHandler.DeleteMenu)
		}
	}
}
