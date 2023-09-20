package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/pflag"
)

var commandMap map[string]func(string) ([]string, error)

func init() {
	commandMap = make(map[string]func(string) ([]string, error))
	commandMap["c"] = countBytes
	commandMap["w"] = countWords
	commandMap["l"] = countLines
	commandMap["m"] = countMultiBytes
}

func EvaluateFlag(fp *pflag.Flag) {
	var results []string
	var err error
	val, ok := commandMap[fp.Shorthand]
	if ok {
		results, err = val(fp.Value.String())
	} else {
		results, err = countAll(fp.Value.String())
	}
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; i < len(results); i++ {
			fmt.Print(results[i] + " ")
		}
		fmt.Println(fp.Value.String())
	}
}

func countBytes(filePath string) ([]string, error) {
	file, _ := os.Open(filePath)
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanBytes)
	var results []string
	var count int = 0
	for sc.Scan() {
		count = count + 1
	}
	results = append(results, strconv.Itoa(count))
	return results, nil
}

func countMultiBytes(filePath string) ([]string, error) {
	file, _ := os.Open(filePath)
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanRunes)
	var count int = 0
	var results []string
	for sc.Scan() {
		count = count + 1
	}
	results = append(results, strconv.Itoa(count))
	return results, nil
}

func countWords(filePath string) ([]string, error) {
	file, _ := os.Open(filePath)
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanWords)
	var count int = 0
	var results []string
	for sc.Scan() {
		count = count + 1
	}
	results = append(results, strconv.Itoa(count))
	return results, nil
}

func countLines(filePath string) ([]string, error) {
	file, _ := os.Open(filePath)
	sc := bufio.NewScanner(file)
	sc.Split(bufio.ScanLines)
	var results []string
	var count int = 0
	for sc.Scan() {
		count = count + 1
	}
	results = append(results, strconv.Itoa(count))
	return results, nil
}

func countAll(filePath string) ([]string, error) {
	var results []string

	countBytes, err := countBytes(filePath)
	if err != nil {
		return nil, err
	} else {
		results = append(results, countBytes[0])
	}

	countWords, err := countWords(filePath)
	if err != nil {
		return nil, err
	} else {
		results = append(results, countWords[0])
	}

	countLines, err := countLines(filePath)
	if err != nil {
		return nil, err
	} else {
		results = append(results, countLines[0])
	}
	return results, nil
}
