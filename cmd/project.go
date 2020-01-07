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
		opt := &project.Opt{
			TemplateName: templateName,
			PackageName:  packageName,
		}
		if err := project.New(args[0], opt); err != nil {
			panic(err)
		}
	},
}

var projectGenCmd = &cobra.Command{
	Short: "刷新项目",
	Use:   "gen <project>",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		opt := &project.Opt{
			TemplateName: templateName,
			PackageName:  packageName,
		}
		if err := project.Gen(args[0], opt); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
	projectCmd.AddCommand(projectNewCmd)
	projectCmd.AddCommand(projectGenCmd)
}
