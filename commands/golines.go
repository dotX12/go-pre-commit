package commands

import (
	"github.com/spf13/cobra"
)

var goLines = &cobra.Command{
	Use:   "golines [files...]",
	Short: "Format your Go code with golines",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runTool("golines", append([]string{"-w"}, dirNames(args)...))
	},
}

func init() {
	rootCmd.AddCommand(goLines)
}
