package cmd

import (
	"github.com/cheetah-fun-gs/legoctl/internal/command/template"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Short: "开发者命令：模板管理",
	Use:   "template",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var templateNewCmd = &cobra.Command{
	Short: "创建模板",
	Use:   "new <project>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := template.New(args[0], templateName); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.AddCommand(templateNewCmd)
}
