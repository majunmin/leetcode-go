package leetcode_0128

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked
func longestConsecutive(nums []int) int {
	numsSet := make(map[int]bool)
	for _, num := range nums {
		numsSet[num] = true
	}
	//
	var result int
	for _, num := range nums {
		// 寻找左边界
		if !numsSet[num-1] {
			continue
		}
		// num 是左边界
		//统计以num为左边界的区间长度
		curLen := 1
		item := num + 1
		for numsSet[item] {
			curLen++
			item++
		}
		result = max(result, curLen)
	}
	return result
}

func solution1(nums []int) int {
	// 记录 当前数字所在的连续子序列的长度
	boundMap := make(map[int]int)
	//for _, num := range nums {
	//	boundMap[num] = 1
	//}
	// process
	var result int
	var curLen int
	for _, num := range nums {
		if _, exist := boundMap[num]; exist {
			continue
		}
		left, right := boundMap[num-1], boundMap[num+1]
		curLen = left + right + 1
		result = max(result, curLen)
		// 将num 加入 boundMap， 该值已经遍历过,可以为任意值
		boundMap[num] = -1
		//更新前区间的左边界和右边的的区间长度
		boundMap[num-1] = curLen
		boundMap[num+1] = curLen
	}
	return result
}
