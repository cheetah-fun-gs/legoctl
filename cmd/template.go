package cmd

import (
	"github.com/cheetah-fun-gs/legoctl/internal/command/template"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Short: "开发者命令：模板管理",
	Use:   "template",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var templateNewCmd = &cobra.Command{
	Short: "创建模板",
	Use:   "new <project>",
	Run: func(cmd *cobra.Command, args []string) {
		template.New(projectPath, templateName)
	},
}

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.AddCommand(templateNewCmd)
}
