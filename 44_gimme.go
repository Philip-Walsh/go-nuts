package main

func Gimme(array [3]int) int {
	if (array[0] > array[1] || array[0] > array[2]) && (array[0] < array[1] || array[0] < array[2]) {
		return 0
	}
	if (array[1] > array[0] || array[1] > array[2]) && (array[1] < array[0] || array[1] < array[2]) {
		return 1
	}
	return 2
}
