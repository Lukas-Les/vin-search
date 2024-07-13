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
	flag.StringVar(&vinRegex, "v", "\\b[A-HJ-NPR-Z0-9]{17}\\b", "regular expression pattern for a VIN")
	flag.IntVar(&maxResult, "max", -1, "max number of vin numbers to extract")
	flag.Parse()

	var outPut []string

	if targetFilePath != "" {
		f, err := os.ReadFile(targetFilePath)
		check(err)
		outPut = []string{
			formatOutput(
				targetFilePath,
				search(string(f), vinRegex, maxResult),
			)}
	}
	if outputFilePath != "" {
		writeOutput(outPut, outputFilePath)
	} else {
		for _, line := range outPut {
			fmt.Println(line)
		}
	}
}

func search(content string, vinRegex string, maxResult int) []string {
	re := regexp.MustCompile(vinRegex)
	var result []string = re.FindAllString(content, maxResult)
	return result
}
