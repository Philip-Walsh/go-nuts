package main

import (
  "fmt"
  "strings"
)

func ToCsvText(array [][]int) (result string) {
      for i, line := range array {
        result += strings.Trim(strings.Replace(fmt.Sprint(line), " ", ",", -1), "[]")
        if i < len(array)-1 {
          result += "\n"
        }
      }
  return result
}
