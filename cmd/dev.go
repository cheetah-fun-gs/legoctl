package cmd

import (
	"github.com/cheetah-fun-gs/goso-cli/internal/dev"
	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Short: "开发者命令",
	Use:   "dev",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var devTplCmd = &cobra.Command{
	Short: "重新构建模板文件",
	Use:   "tpl",
	Run: func(cmd *cobra.Command, args []string) {
		dev.Tpl()
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
	devCmd.AddCommand(devTplCmd)
}
