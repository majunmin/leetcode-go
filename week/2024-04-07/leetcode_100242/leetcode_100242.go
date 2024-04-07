package leetcode_100242

import "strings"

// https://leetcode.cn/contest/weekly-contest-392/problems/lexicographically-smallest-string-after-operations-with-constraint/
func getSmallestString(s string, k int) string {
	if k == 0 {
		return s
	}
	var sbuilder strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'a' {
			sbuilder.WriteByte(c)
			continue
		}
		//是否可以转化为a
		if int(c-'a') <= k || 26-int(c-'a') <= k {
			mina := min(int(c-'a'-0), int(26-(c-'a')))
			k -= mina
			sbuilder.WriteByte('a')
		} else {
			sbuilder.WriteByte(c - uint8(k))
		}
	}
	return sbuilder.String()
}

func distance(s1, s2 string) int {
	var res int
	for i := 0; i < len(s1); i++ {
		c1, c2 := s1[i]-'a', s2[i]-'a'
		res += (int(c1-c2) + 26) % 26
	}
	return res
}
