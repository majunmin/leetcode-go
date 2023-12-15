package leetcode_0541

// https://leetcode.cn/problems/reverse-string-ii/
func reverseStr(s string, k int) string {
	sRaw := []byte(s)

	for i := 0; i < len(sRaw); i += k {
		reverseStrInternal(sRaw, i, i+k)
	}

	return string(sRaw)
}

func reverseStrInternal(raw []byte, begin, end int) {
	end = min(len(raw)-1, end)
	for begin < end {
		raw[begin], raw[end] = raw[end], raw[begin]
		begin++
		end--
	}
}
