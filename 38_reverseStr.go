package main

import (
  "unicode"
)
func ReverseLetters(s string) (res string) {
  for _, r := range s {
    if unicode.IsLetter(r) {
      res = string(r) + res
    }
  }
  return res
}
