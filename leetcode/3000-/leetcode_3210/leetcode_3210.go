package leetcode_3210

import "strings"

// https://leetcode.cn/problems/find-the-encrypted-string/
func getEncryptedString(s string, k int) string {
	// 相当于 将字符串 s 向右循环右移 len(s) - k  || 循环左移  k 位
	n := len(s)
	k = n - k%n
	var sb strings.Builder
	for i := k; i < k+n; i++ {
		sb.WriteByte(s[i%n])
	}
	return sb.String()
}
