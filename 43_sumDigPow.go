package main

import (
  "strconv"
)

func SumDigPow(a, b uint64) (result []uint64) {
  for i := a; i <= b; i++ {
    if isEureka(i) {
      result = append(result, i)
    }
  }
  return result
}

func isEureka(num uint64) bool {
  numStr := strconv.FormatUint(num, 10)
  var sum uint64 = 0
  for i, c := range numStr {
    digit, err := strconv.ParseUint(string(c), 10, 64)
    if err != nil {
      return false
    }
    sum += pow(digit, uint64(i+1))
  }
  return sum == num
}


func pow(base, exp uint64) uint64 {
  result := uint64(1)
  for exp > 0 {
    if exp%2 == 1 {
      result *= base
    }
    base *= base
    exp /= 2
  }
  return result
}
