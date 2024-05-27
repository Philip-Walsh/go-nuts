package main

import (
	"strings"
	"unicode"
)

func solve(str string) string {
	var lowerCount = 0
	var upperCount = 0
	for _, s := range str {
		if unicode.IsLower(s) {
			lowerCount += 1
		} else {
			upperCount += 1
		}
	}
	if upperCount > lowerCount {
		return strings.ToUpper(str)
	}
	return strings.ToLower(str)

}
