package main

import (
  "strings"
)

func High(str string) (res string) {
    var topScore int = 0
    for _, word := range strings.Fields(str) {
      var score int = 0
      for _, char := range word {
        score += int(char) - 96
      }
      if score > topScore {
        res = word
        topScore = score
      }
    }
  return res
}