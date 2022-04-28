// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-28 09:00
package leetcode_0088

//https://leetcode-cn.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {
	// checkParam
	//
	l := len(nums1)
	idx := l - 1
	for ; idx >= 0; idx-- {
		if m == 0 || n == 0 {
			break
		}
		if nums1[m-1] >= nums2[n-1] {
			nums1[idx] = nums1[m-1]
			m -= 1
		} else {
			nums1[idx] = nums2[n-1]
			n -= 1
		}
	}

	if n == 0 {
		return
	}
	// 此时 idx == n-1
	for idx >= 0 {
		nums1[idx] = nums2[idx]
		idx--
	}
}
