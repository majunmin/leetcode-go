package leetcode_0205

func isIsomorphic(s string, t string) bool {
	// 遍历字符串，记录每一个字符最后一次出现的位置.
	// s t 必须一致.
	if len(s) != len(t) {
		return false
	}
	var (
		preS = make(map[byte]int)
		preT = make(map[byte]int)
	)
	for i := 0; i < len(s); i++ {
		if preS[s[i]] != preT[t[i]] {
			return false
		}
		// 避开 0，(不存在是 0)
		// "aa"  "ab"
		preS[s[i]] = i + 1
		preT[t[i]] = i + 1
	}
	return true
}
