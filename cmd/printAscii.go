package cmd

import (
	"fmt"
	"strings"
)

func PrintAsciiArt(words []string, asciiArt map[rune][]string) {
	for _, word := range words {
		PrintWordInAsciiArt(word, asciiArt)
	}
}

func PrintWordInAsciiArt(word string, asciiArt map[rune][]string) {
	var numLines int
	for _, lines := range asciiArt {
		numLines = len(lines)
		break
	}
	if word == "" {
		numLines = 1
	}

	for i := 0; i < numLines; i++ {
		line := ""
		for _, char := range word {
			if lines, ok := asciiArt[char]; ok {
				line += lines[i]
			} else {
				line += " "
			}
		}
		fmt.Println(line)
	}
}

func WithNewLines(str string) string {
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case '\\':
			if i+1 < len(str) {
				nextChar := str[i+1]
				switch nextChar {
				case 'n':
					builder.WriteByte('\n')
				case '\\':
					builder.WriteByte('\\')
				default:
					builder.WriteByte('\\')
					builder.WriteByte(nextChar)
				}
				i++
			} else {
				builder.WriteByte('\\')
			}
		default:
			builder.WriteByte(str[i])
		}
	}
	return builder.String()
}
