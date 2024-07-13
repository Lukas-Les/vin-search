package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var targetFilePath string
	var outputFilePath string
	var vinRegex string
	var maxResult int
	flag.StringVar(&targetFilePath, "f", "", "location of the target file")
	flag.StringVar(&outputFilePath, "o", "", "location of the output file")
	flag.StringVar(&vinRegex, "v", "(?i)[A-HJ-NPR-Z0-9]{17}", "regular expression pattern for a VIN")
	flag.IntVar(&maxResult, "max", -1, "max number of vin numbers to extract")
	flag.Parse()

	var output []string

	if targetFilePath != "" {
		f, err := os.ReadFile(targetFilePath)
		check(err)
		output = []string{
			formatOutput(
				targetFilePath,
				search(string(f), vinRegex, maxResult),
			)}
	}
	if outputFilePath != "" {
		writeOutput(output, outputFilePath)
	} else {
		for _, line := range output {
			fmt.Println(line)
		}
	}
}

func search(content string, vinRegex string, maxResult int) []string {
	re := regexp.MustCompile(vinRegex)
	return re.FindAllString(content, maxResult)
}
