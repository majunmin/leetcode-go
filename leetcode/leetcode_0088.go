package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	// param check
	if n == 0 {
		return
	}
	idx := len(nums1) - 1
	i1, i2 := m-1, n-1
	for i2 > -1 {
		if nums1[i1] < nums2[i2] {
			nums1[idx] = nums1[i1]
			i1--
		} else {
			nums1[idx] = nums2[i2]
			i2--
		}
	}
}
