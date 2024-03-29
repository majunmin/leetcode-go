package leetcode_0049

import "strings"

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string, len(strs))
	// 分组
	for _, s := range strs {
		seqStr := convert2SequenceStr(s)
		resultMap[seqStr] = append(resultMap[seqStr], s)
	}

	// build result
	var result [][]string
	for _, item := range resultMap {
		result = append(result, item)
	}
	return result
}

func convert2SequenceStr(s string) string {
	cnts := [26]int{}
	for i := range s {
		cnts[s[i]-'a']++
	}
	var sb strings.Builder
	for i, item := range cnts {
		sb.WriteByte(byte(item + 'a'))
		sb.WriteByte(byte(i + '0'))
	}
	return sb.String()
}
