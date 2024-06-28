package leetcode_2734

// https://leetcode.cn/problems/lexicographically-smallest-string-after-substring-operation/
func smallestString(s string) string {
	// param check
	if len(s) == 0 {
		return ""
	}
	var (
		sRaw = []byte(s)
		i    int
	)
	if s[0] == 'a' {
		i = 1
	}
	for ; i < len(s); i++ {
		if s[i] == 'a' {
			break
		}
		sRaw[i] = sRaw[i] - 1
	}
	return string(sRaw)
}
