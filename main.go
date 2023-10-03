package main

import (
	commands "ccwc/cmd"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func init() {
	rootCmd = &cobra.Command{
		Use:   "ccwc",
		Short: "Run custom wc cli tool",
		Run: func(cmd *cobra.Command, args []string) {
			var results []string
			if len(args) == 0 {
				var fileContent []byte
				fileContent, err := io.ReadAll(os.Stdin)
				if err != nil {
					fmt.Println("Error:\n", err)
					os.Exit(1)
				}
				results = commands.Evaluate(cmd, fileContent)
				for i := 0; i < len(results); i++ {
					fmt.Print(results[i] + " ")
				}
				fmt.Println()
			} else {
				filePath := args[0]
				fileContent, err := os.ReadFile(filePath)
				if err != nil {
					fmt.Println("Error:\n", err)
					os.Exit(2)
				}
				results = commands.Evaluate(cmd, fileContent)
				for i := 0; i < len(results); i++ {
					fmt.Print(results[i] + " ")
				}
				fmt.Println(filePath)
			}
		},
	}
	rootCmd.Flags().BoolP("bytes", "c", false, "Count number of bytes in the file")
	rootCmd.Flags().BoolP("lines", "l", false, "Count number of lines in the file")
	rootCmd.Flags().BoolP("words", "w", false, "Count number of words in the file")
	rootCmd.Flags().BoolP("multibytes", "m", false, "Count number of multibytes in the file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
