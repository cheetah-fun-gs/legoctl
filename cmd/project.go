package cmd

import (
	"github.com/cheetah-fun-gs/legoctl/internal/command/project"
	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Short: "项目管理",
	Use:   "project",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var projectNewCmd = &cobra.Command{
	Short: "创建项目",
	Use:   "new <project>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := project.New(args[0], templateName); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectNewCmd)
}
