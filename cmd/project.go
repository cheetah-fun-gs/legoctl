package cmd

import (
	"github.com/cheetah-fun-gs/legoctl/internal/command/gen"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "gen <project>",
	Short: "为lego项目生成代码",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		gen.Gen(args[0])
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
