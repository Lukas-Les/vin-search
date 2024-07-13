package main

import (
	"strings"
	"testing"
)

func TestSearch(t *testing.T) {
	var lines = []string{"test VF34C5FWC8S014204 eof", "test vf34c5fwc8s014204 eof"}
	var want = []string{"VF34C5FWC8S014204"}
	for _, line := range lines {
		var result []string = search(line, "(?i)\\b[A-HJ-NPR-Z0-9]{17}\\b", -1)
		if len(result) == 0 {
			t.Fatalf("No VIN found in %s", line)
		}
		strResult := strings.ToUpper(result[0])
		if strResult != want[0] {
			t.Fatalf("%s is not equal to %s", result, want)
		}
	}
}
