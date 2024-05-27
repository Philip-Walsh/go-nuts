package main

import (
	"strings"
)

func Points(games []string) (points int) {
	for _, i := range games {
		var outcome = strings.Split(i, ":")
		if outcome[0] > outcome[1] {
			points += 3
		} else if outcome[0] == outcome[1] {
			points += 1
		}
	}
	return points
}
