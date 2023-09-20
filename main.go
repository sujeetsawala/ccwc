package main

import (
	commands "ccwc/cmd"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	var flag string
	rootCmd = &cobra.Command{
		Use:   "ccwc",
		Short: "Run custom wc cli tool",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Flags().Visit(commands.EvaluateFlag)
		},
	}
	rootCmd.Flags().StringVarP(&flag, "bytes", "c", "a", "Count number of bytes in the file")
	rootCmd.Flags().StringVarP(&flag, "lines", "l", "a", "Count number of lines in the file")
	rootCmd.Flags().StringVarP(&flag, "words", "w", "a", "Count number of words in the file")
	rootCmd.Flags().StringVarP(&flag, "multibytes", "m", "a", "Count number of multibytes in the file")
}

func main() {
	rootCmd.Execute()
}
