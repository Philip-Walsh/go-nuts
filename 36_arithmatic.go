package main

func Arithmetic(a int, b int, operator string) int{
  switch operator {
    case "add":
      return a + b
    case "subtract":
      return a - b
    case "divide":
      return a / b
  }
  return a * b
}