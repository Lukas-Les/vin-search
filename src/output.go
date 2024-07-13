package main

import (
	"fmt"
	"os"
)

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

func formatOutput(sourceFileName string, results []string) string {
	var output string = sourceFileName + "->"
	if len(results) == 0 {
		return output
	}
	for _, item := range results {
		output += item + " "
	}
	return output
}
