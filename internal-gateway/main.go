package main

import "github.com/gly-hub/go-admin/internal-gateway/cmd"

// @title go-admin API
// @version 1.0.0
// @description 基于RpcX + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
