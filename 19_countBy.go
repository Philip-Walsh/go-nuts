package main

func CountBy(x, n int) (res []int) {
	for i := 1; i <= n; i++ {
		res = append(res, i*x)
	}
	return res
}
