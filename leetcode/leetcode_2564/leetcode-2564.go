package leetcode_2564

import "strings"

var (
	Unknown = []int{-1, -1}
)

// val ^ first = second => val = first ^ second
func substringXorQueries(s string, queries [][]int) [][]int {
	// param check
	if len(s) == 0 || len(queries) == 0 {
		return nil
	}

	memo := make(map[int][]int)
	if i := strings.Index(s, "0"); i >= 0 {
		memo[0] = []int{i, i}
	}
	// 枚举 s 中可以组成的所有的数,
	for l, c := range s {
		if c == '0' {
			continue
		}
		// pow(2, 30) > pow(10, 9)
		for r, x := l, 0; r < l+30 && r < len(s); r++ {
			x = x<<1 | int(s[r]-'0')
			if _, exist := memo[x]; !exist {
				memo[x] = []int{l, r}
			}
		}
	}
	result := make([][]int, 0, len(queries))
	for _, item := range queries {
		val := item[0] ^ item[1]
		if v, exist := memo[val]; exist {
			result = append(result, v)
		} else {
			result = append(result, Unknown)
		}
	}
	return result
}
