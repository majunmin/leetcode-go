package leetcode_3233

import "math"

// 符合条件的 是质数的平方， 用 埃氏筛. 筛选质数.
// https://leetcode.cn/problems/find-the-count-of-numbers-which-are-not-special/
func nonSpecialCount(l int, r int) int {
	right := int(math.Sqrt(float64(r)))
	isPrimary := make([]bool, right+1)
	for i := 0; i <= right; i++ {
		isPrimary[i] = true
	}
	result := r - l + 1
	for i := 2; i <= right; i++ {
		if isPrimary[i] {
			if i*i >= l && i*i <= r {
				result--
			}
			for j := i + i; j <= right; j += i {
				isPrimary[j] = false
			}
		}
	}
	return result
}
