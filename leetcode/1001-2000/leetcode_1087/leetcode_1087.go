package leetcode_1087

import (
	"strings"
)

// https://leetcode.ca/2018-11-21-1087-Brace-Expansion/
func expand(s string) []string {
	items := convert(s)
	result := make([]string, 0)
	dfs(items, 0, "", &result)
	return result
}

func dfs(items [][]string, level int, path string, result *[]string) {
	if level == len(items) {
		*result = append(*result, path)
		return
	}
	for _, item := range items[level] {
		dfs(items, level+1, path+item, result)
	}

}

func convert(s string) [][]string {
	if len(s) == 0 {
		return nil
	}
	result := make([][]string, 0)
	prev := ""
	for i := 0; i < len(s); i++ {
		j := i
		if s[i] == '{' {
			// find }
			if len(prev) > 0 {
				result = append(result, []string{})
				result[len(result)-1] = append(result[len(result)-1], prev)
				prev = ""
			}
			j = i
			for ; j < len(s); j++ {
				if s[j] == '}' {
					break
				}
			}
			result = append(result, strings.Split(s[i+1:j], ","))
			i = j
		} else {
			prev += string(s[i])
		}
	}
	return result
}
