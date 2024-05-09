package leetcode_0242

import "sort"

// https://leetcode-cn.com/problems/valid-anagram/description/
func isAnagram(s string, t string) bool {
	return arraySolution(s, t)
}

func hashSolution(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	m := make(map[rune]int, 26)

	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		m[r]--
		if m[r] < 0 {
			return false
		}
	}
	return true

}

// 由于 stringstr 仅由 26 个小写字母组成, 可以用一个数组 枚举出所有值
func arraySolution(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
		arr[t[i]-'a']--
	}
	//
	for i := 0; i < 26; i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

// sortSolution
func sortSolution(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	bytes1, bytes2 := []byte(s), []byte(t)
	sort.Slice(bytes1, func(i, j int) bool { return bytes1[i] < bytes1[j] })
	sort.Slice(bytes2, func(i, j int) bool { return bytes2[i] < bytes2[j] })
	return string(bytes1) == string(bytes2)
}
