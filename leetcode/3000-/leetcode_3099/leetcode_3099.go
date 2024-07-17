package leetcode_3099

// https://leetcode.cn/problems/harshad-number/
func sumOfTheDigitsOfHarshadNumber(x int) int {
	// param check
	if x == 0 {
		return -1
	}

	var (
		sum int
		y   = x
	)
	for y > 0 {
		sum += y % 10
		y /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
