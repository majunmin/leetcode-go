package offer_058

func reverseLeftWords(s string, n int) string {

	if n == 0 || n > len(s) {
		return s
	}
	return solution2(s, n)
}

// 根据n 将 原字符串 C 分解为 两个 字符串 A  B
// 反转A => A'  反转 B => B'
// 反转 A'B' =>  结果
func solution2(s string, n int) string {
	runes := []rune(s)
	reverse(runes[:n])
	reverse(runes[n:])
	reverse(runes)
	return string(runes)
}

func reverse(rs []rune) {
	i, j := 0, len(rs)-1
	for i < j {
		rs[i], rs[j] = rs[j], rs[i]
		i++
		j--
	}
}

func solution1(s string, n int) string {
	//return s[n:] + s[:n]
	for i := 0; i < n; i++ {
		s += string(s[i])
	}
	return s[n:]
}
