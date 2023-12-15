package leetcode_0344

// https://leetcode.cn/problems/reverse-string/
func reverseString(s []byte) {
	if len(s) < 2 {
		return
	}
	left, right := 0, len(s)-1
	//
	//for left < right {
	//	s[left], s[right] = s[right], s[left]
	//	left++
	//	right--
	//}
	// 交换两个 int 类型
	for left < right {
		s[left] ^= s[right]
		s[right] ^= s[left]
		s[left] ^= s[right]
	}
}
