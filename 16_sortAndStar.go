//   var res []int
// 	for _, val := range arr {
//     var sum int = 0
//     for _, char := range []byte(val) {
//       sum += int(char)
//     }
//     res = append(res, sum)
// 	}
//   var smallest = res[0]
//   var index = 0
//   for i := 1; i < len(arr); i++ {
//     if res[i] < smallest {
//       smallest := res[i]
//       index := i
//     }
//   }
// 	return strings.Join(strings.Split(arr[index], ""), "***")

package main

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	first := arr[0]
	return strings.Join(strings.Split(first, ""), "***")
}
