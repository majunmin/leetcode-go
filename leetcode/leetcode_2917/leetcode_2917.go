package leetcode_2917

// https://leetcode.cn/problems/find-the-k-or-of-an-array/
func findKOr(nums []int, k int) int {
	var result int
	for i := 0; i < 31; i++ {
		var cnt int
		for _, num := range nums {
			if num>>i&1 == 1 {
				cnt++
			}
		}
		if cnt >= k {
			result |= 1 << i
		}
	}
	return result
}

func solution1(nums []int, k int) int {
	// param check
	if k > len(nums) {
		return 0
	}
	//
	cnts := make([]int, 32)
	for _, num := range nums {
		for i := 0; i < 32; i++ {
			if num>>i&1 == 1 {
				cnts[i]++
			}
		}
	}
	var result int
	for i, cnt := range cnts {
		if cnt >= k {
			result += 1 << i
		}
	}
	return result
}
