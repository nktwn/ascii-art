package cmd

import (
	"ascii-art/validation"
	"bufio"
	"fmt"
	"os"
)

func ReadAsciiArt(filename string) (map[rune][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if !validation.IsFileCorrect() {
		fmt.Println("Error reading file: File has been changed!")
		os.Exit(1)
	}

	asciiArt := make(map[rune][]string)
	scanner := bufio.NewScanner(file)

	var artLines []string
	var currentChar rune = ' '

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(artLines) > 0 {
				asciiArt[currentChar] = artLines
				artLines = []string{}
				currentChar++
			}
			continue
		}
		artLines = append(artLines, line)
	}
	if len(artLines) > 0 {
		asciiArt[currentChar] = artLines
	}

	return asciiArt, scanner.Err()
}
