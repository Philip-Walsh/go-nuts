package main

func IsTriangle(a, b, c int) bool {
  return a > 0 && b > 0 && c > 0 && 
    (a + b > c) &&
    (a + c > b) &&
    (b + c > a)
}