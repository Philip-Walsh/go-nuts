package main

import "math"

func StringToNumber(str string) int {
	var result = 0
	var isNeg = false
	if str[0:1] == "-" {
		isNeg = true
		str = str[1:]
	}
	for i, num := range str {
		var mult = int(math.Pow(float64(10), float64((len(str) - i - 1))))
		result += mult * (int(num) - 48)
	}
	if isNeg {
		result *= -1
	}
	return result
}
