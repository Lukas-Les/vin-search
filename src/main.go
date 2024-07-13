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

	f, err := os.ReadFile(targetFilePath)
	check(err)
	content := string(f)

	var result []string = search(content, vinRegex, maxResult)

	if outputFilePath != "" {
		writeOutput(result, outputFilePath)
	} else {
		fmt.Println(result)
	}
}

func writeOutput(result []string, outputFilePath string) {
	outFile, err := os.Create(outputFilePath)
	check(err)
	defer outFile.Close()
	for _, item := range result {
		_, err := outFile.WriteString(item + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func search(content string, vinRegex string, maxResult int) []string {
	re := regexp.MustCompile(vinRegex)
	var result []string = re.FindAllString(content, maxResult)
	return result
}
