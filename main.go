package main

import (
	"ascii-art/cmd"
	"ascii-art/validation"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: it only has to be one argument")
		os.Exit(1)
	}
	if !validation.IsASCII(os.Args[1]) {
		fmt.Println("Usage: it should only be in ascii format!")
		os.Exit(1)
	}
	asciiArt, err := cmd.ReadAsciiArt("banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading ASCII art file:", err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	input := os.Args[1]
	if input == "" {
		return
	}

	words := strings.Split(cmd.WithNewLines(input), "\n")
	if !validation.OnlyHasNewLines(words) {
		words = words[1:]
	}
	cmd.PrintAsciiArt(words, asciiArt)
}
