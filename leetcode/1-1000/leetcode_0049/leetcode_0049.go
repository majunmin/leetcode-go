package leetcode_0049

import (
	"sort"
	"strconv"
	"strings"
)

//https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	return arrSolution(strs)
}

func arrSolution(strs []string) [][]string {
	m := make(map[string][]string)

	for _, str := range strs {
		arr := make([]int, 26)
		for i := 0; i < len(str); i++ {
			arr[str[i]-'a']++
		}
		var sb strings.Builder
		for i := range arr {
			if arr[i] > 0 {
				sb.WriteByte(byte(i))
				sb.WriteString(strconv.Itoa(arr[i]))
			}
		}
		m[sb.String()] = append(m[sb.String()], str)
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 时间复杂度: O(nklogk) : k 字符串长度， n 字符串个数
// 空间复杂度: O(nk) 使用一个额外的 hash表,n 字符串个数, k字符串长度
func sortSolution(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		m[string(bs)] = append(m[string(bs)], str)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
