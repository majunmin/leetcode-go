package leetcode_0415

// https://leetcode.cn/problems/add-strings/
func addStrings(num1 string, num2 string) string {
	// 双指针.
	var (
		m, n   = len(num1), len(num2)
		i, j   = m - 1, n - 1
		result = make([]byte, 0, max(m, n))
		// 进位
		carry = 0
	)
	for i >= 0 && j >= 0 {
		n1 := int(num1[i] - '0')
		n2 := int(num2[j] - '0')
		cur := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		result = append([]byte{byte(cur + '0')}, result...)
		i--
		j--
	}
	for i >= 0 {
		cur := (int(num1[i]-'0') + carry) % 10
		carry = (int(num1[i]-'0') + carry) / 10
		result = append([]byte{byte(cur + '0')}, result...)
		i--
	}
	for j >= 0 {
		cur := (int(num2[j]-'0') + carry) % 10
		carry = (int(num2[j]-'0') + carry) / 10
		result = append([]byte{byte(cur + '0')}, result...)
		j--
	}
	if carry == 1 {
		result = append([]byte{'1'}, result...)
	}
	return string(result)
}
