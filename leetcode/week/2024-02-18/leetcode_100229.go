package _024_02_18

import "strconv"

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	// 统计 arr1 中 所有的前缀字符串
	has := make(map[string]bool)
	for _, x := range arr1 {
		s := strconv.Itoa(x)
		for i := 1; i < len(s); i++ {
			has[s[:i]] = true
		}
	}
	// 与 arr2 中 字符串进行比较
	var result int
	for _, x := range arr2 {
		s := strconv.Itoa(x)
		for i := 1; i < len(s); i++ {
			if has[s[:i]] {
				result = max(result, i)
			}
		}
	}
	return result
}

func commonPrefixLen(str1, str2 string) int {
	var length int
	minLen := min(len(str1), len(str2))
	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			break
		}
		length++
	}
	return length
}
