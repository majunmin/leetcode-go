package leetcode_2589

import (
	"slices"
)

// https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/
func findMinimumTime(tasks [][]int) int {
	// 思路: 贪心
	// 1. 按照 end 排序.
	// 2.1 第一个任务 优先从后往前消耗.
	// 2.2 第二个任务优先从前往后 使用 与 前面任务 重合的时间点.
	// 2.3 然后第二个任务 在使用从后往前使用的时间点.
	// 3. 后面的任务依次类推.

	slices.SortFunc(tasks, func(a, b []int) int {
		return a[1] - b[1]
	})
	var (
		mx     = tasks[len(tasks)-1][1] + 1
		run    = make([]bool, mx)
		result int
	)
	for _, task := range tasks {
		start, end, d := task[0], task[1], task[2]
		for i := start; i <= end; i++ {
			if run[i] {
				d--
			}
		}
		for i := end; i >= start; i-- {
			if d == 0 {
				break
			}
			if !run[i] {
				run[i] = true
				d--
				result++
			}
		}
	}
	return result
}
