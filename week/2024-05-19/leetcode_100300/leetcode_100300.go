package leetcode_100300

func sumDigitDifferences(nums []int) int64 {
	// 数位个数.
	var (
		result int
		idx    int
	)
	pivot := nums[0]
	for pivot > 0 {
		cntMap := make(map[int]int)
		for _, num := range nums {
			num = num / pow(10, idx)
			cur := num % 10
			cntMap[cur]++
		}
		if len(cntMap) <= 1 {
			result += 0
		} else {
			res := 1
			for _, v := range cntMap {
				res *= v
			}
			result += res
		}
		pivot /= 10
		idx += 1
	}
	return int64(result)
}

func pow(num int, n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return num
	}
	return num * pow(num, n-1)
}
