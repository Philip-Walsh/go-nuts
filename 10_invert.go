package main

func Invert(arr []int) (result []int) {
	for _, val := range arr {
		result = append(result, val*-1)
	}
	return result
}
