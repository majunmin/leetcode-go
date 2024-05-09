package leetcode_1052

// https://leetcode.cn/problems/grumpy-bookstore-owner/
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	var (
		maxConvert int
		convert    int // 窗口内不满意的顾客数
		satisfied  int
		n          = len(customers)
	)
	for i := 0; i < n; i++ {
		if grumpy[i] == 0 {
			satisfied += customers[i]
		}
		// cal result && shrink window
		if i >= minutes {
			maxConvert = max(maxConvert, convert)
			if grumpy[i-minutes] == 1 {
				convert -= customers[i-minutes]
			}
		}
		// add to window
		if grumpy[i] == 1 {
			convert += customers[i]
		}
	}
	return satisfied + maxConvert
}
