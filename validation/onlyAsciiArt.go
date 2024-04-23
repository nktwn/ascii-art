package validation

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

const KNOWNHASH = "ac85e83127e49ec42487f272d9b9db8b"

func IsASCII(s string) bool {
	for _, char := range s {
		if char > 127 {
			return false
		}
	}
	return true
}

func IsFileCorrect() bool {
	fileContent, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return false
	}
	hash := md5.Sum(fileContent)
	computedHash := hex.EncodeToString(hash[:])
	if computedHash == KNOWNHASH {
		return true
	} else {
		return false
	}
}

func OnlyHasNewLines(words []string) bool {
	hasOnlyNewlines := false
	for _, word := range words {
		if word != "" {
			hasOnlyNewlines = true
		}
	}
	return hasOnlyNewlines
}


