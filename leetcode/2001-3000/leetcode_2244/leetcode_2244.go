package leetcode_2244

// https://leetcode.cn/problems/minimum-rounds-to-complete-all-tasks/description/
func minimumRounds(tasks []int) int {
	var (
		// 统计每种任务的 计数
		taskNums = make(map[int]int)
		result   int
	)
	for _, level := range tasks {
		taskNums[level]++
	}
	for _, num := range taskNums {
		if num == 1 {
			return -1
		}
		if num%3 == 0 {
			result += num / 3
		} else {
			result += num/3 + 1
		}
	}
	return result
}
