package internal

import (
	"fmt"
	"os"
	"strings"
)

func WriteOutput(result []string, outputFilePath string) {
	outFile, err := os.Create(outputFilePath)
	Check(err)
	defer outFile.Close()
	for _, item := range result {
		_, err := outFile.WriteString(item + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func FormatOutput(sourceFileName string, results []string) string {
	var output string = sourceFileName + "->"
	output += strings.Join(results, " ")
	return output
}
