// package kata

// func Maps(x []int) []int {
//   var res []int
//   for _, i := range x {
//     res = append(res, i*2)
//   }
//   return res
// }

package kata

func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, v := range x {
		result[i] = v + v
	}
	return result
}
