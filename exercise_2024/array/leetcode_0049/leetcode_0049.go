package leetcode_0049

import "strings"

// https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-interview-150
func groupAnagrams(strs []string) [][]string {
	var (
		wordSet = make(map[string][]string)
		result  = make([][]string, 0)
	)

	for _, str := range strs {
		m := buildStr(str)
		wordSet[m] = append(wordSet[m], str)
	}
	for _, items := range wordSet {
		result = append(result, items)
	}
	return result
}

func buildStr(str string) string {
	cnts := [26]int{}
	for i := range str {
		cnts[str[i]-'a']++
	}
	var sBuilder strings.Builder
	for i, cnt := range cnts {
		if cnt == 0 {
			continue
		}
		if cnt > 0 {
			sBuilder.WriteByte(byte(i) + 'a')
			sBuilder.WriteByte(byte(cnt))
		}
	}
	return sBuilder.String()
}
