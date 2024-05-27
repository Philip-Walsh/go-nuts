package main

func Summation(n int) (result int) {
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}
