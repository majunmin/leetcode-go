package leetcode_2663

// https://leetcode.cn/problems/lexicographically-smallest-beautiful-string/description/
func smallestBeautifulString(s string, k int) string {
	var (
		n      = len(s)
		cRight = byte('a' + k)
		sRaw   = []byte(s)
		i      int
		flag   bool // 是否存在结果
	)
	for i = n - 1; i >= 0; i-- {
		for c := s[i] + 1; c < cRight; c++ {
			if checkOk(sRaw, i, c) {
				sRaw[i] = c
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	// 在 s 的位置上防止 字符c
	// 在 i 的位置后面生成最小字典序的字符串 s[i+1:]
	if !flag {
		return ""
	}
	for j := i + 1; j < n; j++ {
		for c := byte('a'); c < cRight; c++ {
			if checkOk(sRaw, j, c) {
				sRaw[j] = c
				break
			}
		}
	}
	return string(sRaw)
}

// s 在 i的位置上防止字符j 是否是回文串
func checkOk(s []byte, i int, c byte) bool {
	//return  i > 0 && s[i-1] != c || i < len(s)-1 && s[i] == s[i+1] || s[]
	if i >= 1 && s[i-1] == c {
		return false
	}
	if i >= 2 && s[i-2] == c {
		return false
	}
	return true
}
