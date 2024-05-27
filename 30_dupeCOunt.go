package main

import (
	"strings"
	"unicode"
)

func duplicate_count(s1 string) int {
	countMap := make(map[rune]int)

	var s = strings.ToLower(s1)

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			countMap[char]++
		}
	}

	duplicateCount := 0
	for _, count := range countMap {
		if count > 1 {
			duplicateCount++
		}
	}

	return duplicateCount

}
