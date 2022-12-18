package route

import (
	"github.com/gly-hub/go-admin/internal-gateway/internal/route/auth"
	"github.com/gly-hub/go-dandelion/application"
)

func InitRoute(){
	baseRouter := application.HttpServer().Router()
	// auth服务相关路由
	auth.InitAuthRoute(baseRouter)
}
