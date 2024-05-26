package main

import (
	"math"
)

func SquareSum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += int(math.Pow(float64(val), 2))
	}
	return sum
}
