package auth

import (
	"github.com/gly-hub/go-admin/external-gateway/internal/service/auth"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/server/http"
)

func InitAuthRoute(){
	authHandler := auth.AuthController{}
	// 登录登出
	application.HttpServer().RegisterRoute("", []http.Route{
		http.Route{Path: "/login", Method: http.POST, Handler: authHandler.Login}, // 登录
		//http.Route{Path: "/login", Method: http.POST, Handler: nil}, // 登出
	})

	// 权限菜单
	//application.HttpServer().RegisterRoute("", []http.Route{
	//	http.Route{Path: "/login", Method: http.POST, Handler: nil}, // 登录
	//	http.Route{Path: "/login", Method: http.POST, Handler: nil}, // 登出
	//})
}
