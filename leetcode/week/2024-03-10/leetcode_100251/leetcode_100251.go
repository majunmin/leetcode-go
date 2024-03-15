package leetcode_100251

import (
	"slices"
	"strings"
)

// https://leetcode.cn/contest/weekly-contest-388/problems/shortest-uncommon-substring-in-an-array/
func shortestSubstrings(arr []string) []string {
	var (
		globalSubStrSet = make(map[string]int)
		subStrByIdx     = make([][]string, len(arr))
	)

	for i, item := range arr {
		subStrs := findAllSubStr(item)
		slices.SortFunc(subStrs, func(a, b string) int {
			if len(a) == len(b) {
				return strings.Compare(a, b)
			}
			return len(a) - len(b)
		})
		subStrByIdx[i] = subStrs
		for _, sub := range findAllSubStr(item) {
			globalSubStrSet[sub]++
		}
	}
	//
	result := make([]string, 0, len(arr))
	for _, subStrs := range subStrByIdx {
		var flag bool
		for _, substr := range subStrs {
			// 找到仅出现一次的 substr
			if globalSubStrSet[substr] == 1 {
				result = append(result, substr)
				flag = true
				break
			}
		}
		if !flag {
			// 如果没有找到子字符串， 填入""
			result = append(result, "")
		}
	}
	return result
}

func findAllSubStr(str string) []string {
	// 去掉重复的
	set := make(map[string]bool)
	var result []string
	for i := 1; i <= len(str); i++ {
		for j := 0; j <= len(str)-i; j++ {
			sub := str[j : j+i]
			if !set[sub] {
				result = append(result, sub)
				set[sub] = true
			}
		}
	}
	return result
}
