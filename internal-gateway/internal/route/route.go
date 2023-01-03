package route

import (
	_ "github.com/gly-hub/go-admin/internal-gateway/docs"
	"github.com/gly-hub/go-admin/internal-gateway/internal/route/auth"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/server/http"
	routingSwagger "github.com/gly-hub/go-dandelion/swagger"
)

func InitRoute() {
	baseRouter := application.HttpServer().Router()

	if config.GetEnv() != "production" {
		// 注册swagger
		baseRouter.Get("/swagger/*", routingSwagger.WrapHandler)
		http.LogIgnoreResult(`.*?/swagger/.*?`) // 忽略swagger响应值
	}

	// auth服务相关路由
	auth.InitAuthRoute(baseRouter)
}
