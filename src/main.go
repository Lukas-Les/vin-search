package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var targetFilePath string
	var targetDirPath string
	var outputFilePath string
	var vinRegex string
	var maxResult int
	flag.StringVar(&targetFilePath, "f", "", "location of the target file")
	flag.StringVar(&targetDirPath, "d", "", "location of the target directory")
	flag.StringVar(&outputFilePath, "o", "", "location of the output file")
	flag.StringVar(&vinRegex, "v", "(?i)\\b[A-HJ-NPR-Z0-9]{17}\\b", "regular expression pattern for a VIN")
	flag.IntVar(&maxResult, "max", -1, "max number of vin numbers to extract")
	flag.Parse()

	if (targetDirPath == "" && targetFilePath == "") || (targetDirPath != "" && targetFilePath != "") {
		panic("Please provide one of -f <target file> or -d <target dir>, and not both!")
	}

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
	if targetDirPath != "" {
		var targetFilesPaths []string = collectTargetFilesPaths(targetDirPath)
		for _, filePath := range targetFilesPaths {
			f, err := os.ReadFile(filePath)
			check(err)
			output = append(output, formatOutput(filePath, search(string(f), vinRegex, maxResult)))
		}
	}

	if outputFilePath != "" {
		writeOutput(output, outputFilePath)
	} else {
		for _, line := range output {
			fmt.Println(line)
		}
	}
}

func collectTargetFilesPaths(targetDirPath string) []string {
	var filePaths []string

	err := filepath.Walk(targetDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", targetDirPath, err)
	}

	return filePaths
}

func search(content string, vinRegex string, maxResult int) []string {
	re := regexp.MustCompile(vinRegex)
	return re.FindAllString(content, maxResult)
}
