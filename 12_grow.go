package kata

func Grow(arr []int) int {
	var res int = 1
	for _, val := range arr {
		res = res * val
	}
	return res
}
