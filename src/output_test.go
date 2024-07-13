package main

import (
	"testing"
)

func TestFormatOutput(t *testing.T) {
	sourceFileName := "test.txt"
	var lines = []string{"result1", "result2"}
	want := "test.txt->result1 result2"
	result := formatOutput(sourceFileName, lines)
	if result != want {
		t.Fatalf("%s is not equal to %s", result, want)
	}
}
