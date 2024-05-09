package leetcode_1491

import "math"

// https://leetcode.cn/problems/average-salary-excluding-the-minimum-and-maximum-salary/
func average(salary []int) float64 {
	// param check
	if len(salary) < 3 {
		return 0
	}
	var (
		maxSalary   int
		minSalary   = math.MaxInt
		total       = len(salary)
		totalSalary int
	)
	for _, item := range salary {
		maxSalary = max(maxSalary, item)
		minSalary = min(minSalary, item)
		totalSalary += item
	}
	return float64(totalSalary-minSalary-maxSalary) / float64(total-2)
}
