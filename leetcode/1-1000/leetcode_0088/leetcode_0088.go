// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-28 09:00
package leetcode_0088

//https://leetcode-cn.com/problems/merge-sorted-array/
func merge(nums1 []int, m int, nums2 []int, n int) {

	//如果 n == 0, 说明 nums2 已经插入完成, nums1 剩余的就是有序的
	if n == 0 {
		return
	}

	if m == 0 || nums1[m-1] <= nums2[n-1] {
		nums1[m+n-1] = nums2[n-1]
		merge(nums1, m, nums2, n-1)
	} else {
		nums1[m+n-1] = nums1[m-1]
		merge(nums1, m-1, nums2, n)
	}
}

func solution1(nums1 []int, m int, nums2 []int, n int) {
	// checkParam
	//
	l := len(nums1)
	idx := l - 1
	for ; idx >= 0; idx-- {
		if n == 0 {
			break
		}
		if m == 0 || nums1[m-1] <= nums2[n-1] {
			nums1[idx] = nums2[n-1]
			n--
		} else {
			nums1[idx] = nums1[m-1]
			m--
		}
	}
}
