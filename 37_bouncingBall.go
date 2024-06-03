package main

func BouncingBall(h, bounce, window float64) int {
  if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
    return -1
  }
  var counter = 0

  for h > window {
    counter++
    h *= bounce
    if h > window {
      counter++
    }
  }
  return counter
}