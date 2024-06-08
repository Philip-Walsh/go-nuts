package main

func MultiplicationTable(size int) (result [][]int) {
  for i := 1; i <= size; i++ {
    var row []int
    for j := 1; j <= size; j++ {
        row = append(row, i*j)
    }
  result = append(result, row)
  }
  return result
}