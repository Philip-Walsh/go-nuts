package main

import (
  "unicode"
)

func ToAlternatingCase(str string) string {
  var res string
  for _, r := range str {
    if unicode.IsUpper(r) {
      res += string(unicode.ToLower(r))
    } else {
      res += string(unicode.ToUpper(r))
    }
  }
  return res
}