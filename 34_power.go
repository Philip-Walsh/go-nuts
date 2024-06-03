package main

import (
  "math"
)

func PowersOfTwo(n int) (result []uint64) {
  for i := 0; i <= n; i++ {
    result = append(result, uint64(math.Pow(2, float64(i))))
  }
  return result
}