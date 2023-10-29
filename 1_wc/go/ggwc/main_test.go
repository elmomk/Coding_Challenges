package main_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/elmomk/Coding_Challenges/1_wc/go/ggwc/cmd"
)

func TestCountItemsInFile(t *testing.T) {
	// Create a test case with a mocked file content
	fileContent := "line1\nline2\nline3\n"
	file := strings.NewReader(fileContent)
  table := []struct { // Create a table of test cases
    name string
    splitType bufio.SplitFunc
    expectedCount int
  }{
    {"Lines", bufio.ScanLines, 3},
    {"Words", bufio.ScanWords, 3},
    {"Bytes", bufio.ScanBytes, 18},
  }

  
  fileReader := bufio.NewReader(file)
	// Call the function being tested
  for _, tt := range table {
    t.Logf("Testing %s", tt.name)

	count := cmd.CountItemsInFile(fileReader, tt.splitType)

	// Check if the count matches the expected value
	if count != tt.expectedCount {
		t.Errorf("CountItemsInFile returned %d, expected %d", count, tt.expectedCount)
	}
  }
}
