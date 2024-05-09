package leetcode_0228

import "strconv"

// https://leetcode.cn/problems/summary-ranges/
func summaryRanges(nums []int) []string {
	var (
		result = make([]string, 0, len(nums))
		start  = nums[0]
		end    = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			end = nums[i]
		} else {
			addToResult(end, start, &result)
			start, end = nums[i], nums[i]
		}
	}
	addToResult(end, start, &result)
	return result
}

func addToResult(end int, start int, result *[]string) {
	if end == start {
		*result = append(*result, strconv.Itoa(start))
	} else {
		*result = append(*result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}
}
