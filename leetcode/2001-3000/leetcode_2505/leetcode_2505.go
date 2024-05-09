package leetcode_2505

// https://github.com/doocs/leetcode/blob/main/solution/2500-2599/2505.Bitwise%20OR%20of%20All%20Subsequence%20Sums/README.md
func subsequenceSumOr(nums []int) int64 {
	var (
		ans  int64
		cnts = make([]int, 64)
	)
	for _, num := range nums {
		for i := 0; i < 31; i++ {
			if num>>i&1 == 1 {
				cnts[i]++
			}
		}
	}
	for i := 0; i < 63; i++ {
		// 将 ans 第i位置为1
		if cnts[i] > 0 {
			ans |= 1 << i
		}
		cnts[i+1] = cnts[i] / 2
	}
	return ans
}
