package exercise_2024

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	return float64(getKth(nums1, 0, m-1, nums2, 0, n-1, (m+n+1)>>1)+getKth(nums1, 0, m-1, nums2, 0, n-1, (m+n+2)>>1)) / 2.0
}

func getKth(nums1 []int, s1 int, e1 int, nums2 []int, s2 int, e2 int, k int) int {
	len1, len2 := e1-s1+1, e2-s2+1
	if len1 > len2 {
		return getKth(nums2, s2, e2, nums1, s1, e1, k)
	}
	//让 len1 的长度小于 len2，这样就能保证如果有数组空了，一定是 len1
	if len1 == 0 {
		return nums2[s2+k-1]
	}
	// terminate
	if k == 1 {
		return min(nums1[s1], nums2[s2])
	}
	i := s1 + min(len1, k/2) - 1
	j := s2 + min(len2, k/2) - 1
	if nums1[i] < nums2[j] {
		// 排除掉了 i- s1 +1 个数
		return getKth(nums1, i+1, e1, nums2, s2, e2, k-(i-s1+1))
	} else {
		return getKth(nums1, s1, e1, nums2, j+1, e2, k-(j-s2+1))
	}
}
