package leetcode_2332

import "sort"

// https://leetcode.cn/problems/the-latest-time-to-catch-a-bus/
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	// 模拟乘客上车
	var (
		j int
		c int
	)
	for _, t := range buses {
		for c = capacity; c > 0 && passengers[j] <= t; c-- {
			j++
		}
	}
	// 寻找插队时机
	j--
	result := buses[len(buses)-1]
	if c <= 0 {
		result = passengers[j]
	}
	for j >= 0 && result == passengers[j] {
		result--
		j--
	}
	return result
}
