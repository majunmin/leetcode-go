package leetcode_2801

// https://leetcode.cn/problems/faulty-keyboard/?envType=daily-question&envId=2024-04-01
func finalString(s string) string {
	var (
		result = make([]byte, 0, len(s))
		end    = 0
	)

	for i := 0; i < len(s); i++ {
		if s[i] != 'i' {
			result = append(result, s[i])
			end++
			continue
		}
		reverse(result, 0, end)
	}
	return string(result)
}

func reverse(data []byte, l int, r int) {
	for l < r {
		data[l], data[r] = data[r], data[l]
		l++
		r--
	}
}
