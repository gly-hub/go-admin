package auth

import (
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-admin/internal-gateway/internal/middleware"
	"github.com/gly-hub/go-admin/internal-gateway/internal/service/auth"
)

func InitAuthRoute(baseRouter *routing.RouteGroup) {
	authHandler := auth.AuthController{}

	// 登录登出
	baseRouter.Post("/login", authHandler.Login)

	systemRouter := baseRouter.Group("/system")
	systemRouter.Use(middleware.PermissionMiddleware())
	{
		// 获取系统信息
		systemRouter.Get("/app_info", authHandler.GetSystemInfo)
		// 获取菜单树
		systemRouter.Get("/menu/tree", authHandler.GetMenuTree)
	}

	authGroup := baseRouter.Group("/auth")
	{
		// 菜单管理
		menuGroup := authGroup.Group("/menu")
		{
			menuGroup.Post("/search", authHandler.GetMenuList)
			menuGroup.Post("", authHandler.CreateMenu)
			menuGroup.Put("", authHandler.UpdateMenu)
			menuGroup.Delete("/<id>", authHandler.DeleteMenu)
		}

		// 字典管理
		dictGroup := authGroup.Group("/dict")
		{
			// 标签
			dictGroup.Post("/search", authHandler.GetDictList)
			dictGroup.Post("", authHandler.CreateDict)
			dictGroup.Put("", authHandler.UpdateDict)
			dictGroup.Delete("/<id>", authHandler.DeleteDict)
			// 值
			dictGroup.Post("/value/search", authHandler.GetDictValueList)
			dictGroup.Post("/value", authHandler.CreateDictValue)
			dictGroup.Put("/value", authHandler.UpdateDictValue)
			dictGroup.Delete("/value/<id>", authHandler.DeleteDictValue)
		}
	}
}
