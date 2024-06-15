package main

func Tribonacci(signature [3]float64, n int) []float64 {
  result := signature[:]
  for len(result) < n {
    nextValue := result[len(result)-1] + result[len(result)-2] + result[len(result)-3]
    result = append(result, nextValue)
  }
  return result[:n]
}