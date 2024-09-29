package leetcode_3106

// https://leetcode.cn/problems/lexicographically-smallest-string-after-operations-with-constraint/
func getSmallestString(s string, k int) string {
	raw := []byte(s)
	for i := 0; i < len(s) && k > 0; i++ {
		if s[i] == 'a' {
			continue
		}
		dis := min(abs(int(s[i]-'a')), abs(26-int(s[i]-'a')))
		if dis <= k {
			k -= dis
			raw[i] = 'a'
		} else {
			raw[i] = s[i] - byte(k)
			k = 0
		}
	}
	return string(raw)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// distance(s[i],'a') := min(abs(s[i] - 'a'), abs(26 - s[i] + 'a'）)
