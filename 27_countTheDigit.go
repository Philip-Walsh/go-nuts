package main

import (
	"strconv"
	"strings"
)

func NbDig(n int, d int) int {
	var res int = 0
	for i := 0; i <= n; i++ {
		res += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}

	return res
}
