package main

func Evaporator(content float64, evapPerDay int, threshold int) int {
	var counter int = 0
	var full float64 = 100
	var stop float64 = 100 * (float64(threshold) / 100)

	for full >= stop {
		counter++
		full = full - (full * (float64(evapPerDay) / 100))
	}
	return counter
}
