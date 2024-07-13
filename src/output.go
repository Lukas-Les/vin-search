package main

import (
	"fmt"
	"os"
	"strings"
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
	output += strings.Join(results, " ")
	return output
}
