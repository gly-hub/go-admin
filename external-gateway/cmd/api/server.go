/**
* @Author vangogh
* @Description api服务cmd
* @File:  model
* @Datetime 2022/4/20 10:07
**/
package api

import (
	"github.com/gly-hub/go-admin/external-gateway/internal/route"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/spf13/cobra"
)

var (
	env      string
	StartCmd = &cobra.Command{
		Use:          "model",
		Short:        "Start API model",
		Example:      "external-gateway model -e local",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&env, "env", "e", "local", "Env")
}

func setup() {
	// 配置初始化
	config.InitConfig(env)
	// 应用初始化
	application.Init()
	// 路由初始化
	route.InitRoute()
}

func run() error {
	// 启动http服务
	application.HttpServer().Server()
	return nil
}
