package cmd

import (
	"github.com/cheetah-fun-gs/goso-cli/internal/new"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new <project>",
	Short: "创建一个goso项目",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		new.New(args[0])
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
