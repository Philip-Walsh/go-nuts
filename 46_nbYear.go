package main

func NbYear(p0 int, percent float64, aug int, p int) int {
  var population int = p0
  var years int = 0
  for population < p {
    years++
    population += int(float64(population) * (percent/100)) + aug
  }
  return years
}