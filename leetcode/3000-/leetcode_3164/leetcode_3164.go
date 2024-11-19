package leetcode_3164

func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	return enumSolution(nums1, nums2, k)
}

func enumSolution(nums1 []int, nums2 []int, k int) int64 {
	var (
		cnt1 = make(map[int]int)
		cnt2 = make(map[int]int)
		maxN int
	)
	for _, num := range nums1 {
		cnt1[num]++
		maxN = max(maxN, num)
	}
	for _, num := range nums2 {
		cnt2[num*k]++
	}
	var result int
	for n2, count := range cnt2 {
		for n2 <= maxN {
			if cnt1[n2] > 0 {
				result += cnt1[n2] * count
			}
			n2 += n2
		}
	}
	return int64(result)
}

// 超时
func baoliSolution1(nums1 []int, nums2 []int, k int) int64 {
	var result int
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			if n1%(n2*k) == 0 {
				result++
			}
		}
	}
	return int64(result)
}
