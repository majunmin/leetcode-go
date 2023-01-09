package leetcode_0961

// https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/
func repeatedNTimes(nums []int) int {
	cnt := make(map[int]struct{})
	for _, num := range nums {
		if _, exist := cnt[num]; exist {
			return num
		}
		cnt[num] = struct{}{}
	}
	return -1
}
