package leetcode_1343

// https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
func numOfSubarrays(arr []int, k int, threshold int) int {
	// param check
	if len(arr) == 0 || len(arr) < k {
		return 0
	}
	var (
		l, r int
		cnt  int
		sum  int
	)

	for r < k {
		sum += arr[r]
		r++
	}
	if float64(sum)/float64(k) >= float64(threshold) {
		cnt++
	}
	for r < len(arr) {
		sum = sum + arr[r] - arr[l]
		if float64(sum)/float64(k) >= float64(threshold) {
			cnt++
		}
		r++
		l++
	}
	return cnt
}
