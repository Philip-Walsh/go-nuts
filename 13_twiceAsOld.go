package main

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	return abs(dadYearsOld - 2*sonYearsOld)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
