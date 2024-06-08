package main

import (
	"sort"
	"strings"
)

func NthRank(st string, we []int, n int) (winner string) {
	var names = strings.Split(st, ",")
	if len(names) == 0 || n > len(names) {
		return "Not enough participants"
	}
	nameMap := createNameMap(names, we)
	sort.SliceStable(names, func(i, j int) bool {
		return nameMap[names[i]] > nameMap[names[j]]
	})
	return names[n-1]
}

func createNameMap(names []string, we []int) (nameMap map[string]int) {
	nameMap = make(map[string]int, len(we))
	for i := 0; i < len(we); i++ {
		var name = strings.ToUpper(names[i])
		var nameScore = (len(name) + calcNameScore(name)) * we[i]
		nameMap[names[i]] = nameScore
	}
	return nameMap
}

func calcNameScore(name string) (score int) {
	var DECIMAL_OFFSET = 64
	for _, i := range name {
		score += int(i) - DECIMAL_OFFSET
	}
	return score
}

func sortByValue(nameMap map[string]int) map[string]int {
	keys := make([]string, 0, len(nameMap))
	for key := range nameMap {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		// descending
		return nameMap[keys[i]] < nameMap[keys[j]]
	})

	sortedMap := make(map[string]int, len(nameMap))
	for _, key := range keys {
		sortedMap[key] = nameMap[key]
	}

	return sortedMap
}

func sortByKey(nameMap map[string]int) map[string]int {
	keys := make([]string, 0, len(nameMap))
	for key := range nameMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sortedMap := make(map[string]int, len(nameMap))
	for _, key := range keys {
		sortedMap[key] = nameMap[key]
	}

	return sortedMap
}

// func createNameMap(names []string, we []int) (nameMap map[string]int) {
// 	nameMap = make(map[string]int, len(we))
// 	for i := 0; i <= len(we); i++ {
// 		var name = strings.ToUpper(names[i])
// 		var nameScore = len(name) + calcNameScore(name)*we[i]
// 		nameMap[name] = nameScore
// 	}
// 	return nameMap
// }

// func calcNameScore(name string) (score int) {
// 	var DECIMAL_OFFSET = 64 // 96 for upper
// 	for _, i := range name {
// 		score += int(i) - DECIMAL_OFFSET
// 	}
// 	return score
// }
