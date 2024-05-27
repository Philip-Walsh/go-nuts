package main

func century(year int) (res int) {
	res = int(year / 100)
	if year%100 > 0 {
		res += 1
	}
	return res
}
