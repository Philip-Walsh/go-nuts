package main

import (
	"strings"
)

func PartList(arr []string) (result string ){
    for i := 1; i <= len(arr) - 1; i++ {
      result += "(" + strings.Join(arr[:i], " ") +", " + strings.Join(arr[i:], " ") + ")"
    }
  return result
}