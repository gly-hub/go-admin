/**
* @Author vangogh
* @Description 命令行控制模块
* @File:  cobra
* @Datetime 2022/4/20 10:07
**/
package cmd

import (
	"errors"
	"github.com/gly-hub/go-admin/authorize/cmd/api"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "github.com/gly-hub/go-admin/authorize",
	Short: "authorize",
	SilenceUsage:true,
	Long: "authorize",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New(logger.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init(){
	rootCmd.AddCommand(api.StartCmd)
}

func Execute(){
	if err := rootCmd.Execute(); err != nil{
		os.Exit(-1)
	}
}
