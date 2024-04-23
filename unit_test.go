package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TrimTrailingSpacesAndLines(s string) string {
	lines := strings.Split(s, "\n")
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}
	return strings.Join(lines, "\n")
}

func TestMainFunction(t *testing.T) {
	for i := 1; i <= 15; i++ {
		inputFileName := fmt.Sprintf("tests/input/test%d.txt", i)
		outputFileName := fmt.Sprintf("tests/output/test%d.txt", i)

		t.Run(fmt.Sprintf("TestCase%d", i), func(t *testing.T) {
			input, err := os.ReadFile(inputFileName)
			if err != nil {
				t.Fatalf("Error reading input file: %v", err)
			}
			expectedOutput, err := os.ReadFile(outputFileName)
			if err != nil {
				t.Fatalf("Error reading expected output file: %v", err)
			}

			rescueStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = []string{"cmd", string(input)}
			main()

			w.Close()
			os.Stdout = rescueStdout

			var buf bytes.Buffer
			io.Copy(&buf, r)
			result := TrimTrailingSpacesAndLines(buf.String())
			expected := TrimTrailingSpacesAndLines(string(expectedOutput))
			if result != expected {
				t.Errorf("Test failed for %s. Expected:\n%s\nGot:\n%s", inputFileName, expected, result)
			} else {
				t.Logf("Passed Test Case %d", i)
			}
		})
	}
}
