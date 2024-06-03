package main

import (
	"math"
)

func FindNextSquare(sq int64) int64 {
	root := math.Sqrt(float64(sq))
	if root == float64(int(root)) {
		return int64((root + 1) * (root + 1))
	}
	return -1
}
