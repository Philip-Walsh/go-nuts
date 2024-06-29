package main

func TwoSum(numbers []int, target int) (result [2]int ) {
  var length = len(numbers);
  for i, num := range numbers {
    for j := i + 1; j <= length - 1; j++ {
      if num + numbers[j] == target {
        result = [2]int{i, j}
        return result
      }
    }
  }
  return result;
}