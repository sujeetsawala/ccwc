package commands

import (
	"bytes"
	"strconv"

	"github.com/spf13/cobra"
)

func Evaluate(cmd *cobra.Command, fileContent []byte) []string {
	var results []string
	results = count(fileContent)
	//results = append(results, countWords(bufio.NewScanner(reader))...)
	//results = append(results, countMultibytes(bufio.NewScanner(reader))...)

	var filteredResults []string
	var flagSet bool = false
	if cmd.Flags().Changed("bytes") {
		flagSet = true
		filteredResults = append(filteredResults, results[0])
	}
	if cmd.Flags().Changed("words") {
		flagSet = true
		filteredResults = append(filteredResults, results[2])
	}
	if cmd.Flags().Changed("lines") {
		flagSet = true
		filteredResults = append(filteredResults, results[1])
	}
	if cmd.Flags().Changed("multibytes") {
		flagSet = true
		filteredResults = append(filteredResults, results[3])
	}
	if !flagSet {
		return results
	}
	return filteredResults
}

func count(fileContent []byte) []string {
	var lineCount int = 0
	var byteCount int = len(fileContent)
	var wordCount int = len(bytes.Fields(fileContent))
	var charCount int = len(bytes.Runes(fileContent))

	for i := 0; i < len(fileContent); i++ {
		if string(fileContent[i]) == "\n" {
			lineCount++
		}
	}
	return []string{strconv.Itoa(byteCount), strconv.Itoa(lineCount), strconv.Itoa(wordCount), strconv.Itoa(charCount)}
}
