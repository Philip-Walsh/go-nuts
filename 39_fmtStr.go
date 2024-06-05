package main

import (
  "strings"
  "fmt"
)

func MultiTable(number int) string {
  var result []string
	for i := 1; i <= 10; i++ {
		result =  append(result, fmt.Sprintf("%d * %d = %d", i, number, (i * number)))
	}
	return strings.Join(result, "\n")
}