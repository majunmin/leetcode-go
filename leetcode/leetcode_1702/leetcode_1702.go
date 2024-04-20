package leetcode_1702

// https://leetcode.cn/problems/maximum-binary-string-after-change/?envType=daily-question&envId=2024-04-10
func maximumBinaryString(binary string) string {
	// 滑动窗口
	//

	var (
		data = []byte(binary)
		l, r = 0, 0
	)
	for ; r < len(data); r++ {
		if data[r] == '1' {
			if l == r {
				l++
			}
			continue
		}
		// binary[r] == '0'
		if l == r {
			continue
		}
		data[l] = '1'
		data[r] = '1'
		l++
		data[l] = '0'
	}
	return string(data)
}
