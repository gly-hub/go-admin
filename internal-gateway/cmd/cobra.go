package cmd

import (
	"errors"
	"github.com/gly-hub/go-admin/internal-gateway/cmd/api"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "github.com/gly-hub/go-admin/internal-gateway",
	Short:        "internal-gateway",
	SilenceUsage: true,
	Long:         "internal-gateway",
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

func init() {
	rootCmd.AddCommand(api.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
