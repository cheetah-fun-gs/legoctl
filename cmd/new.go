package cmd

import "github.com/spf13/cobra"

var newCmd = &cobra.Command{
	Use:   "new <project>",
	Short: "创建一个goso项目",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
