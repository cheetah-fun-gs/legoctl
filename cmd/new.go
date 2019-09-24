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

var (
	newProtocols []string
	isWithGate   bool
)

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.LocalFlags().StringArrayVarP(&newProtocols, "protocol", "p", []string{"http"}, "支持的协议：http、websocket、quic")
	newCmd.LocalFlags().BoolVarP(&isWithGate, "with-gate", "g", true, "是否需要网关服务")
}
