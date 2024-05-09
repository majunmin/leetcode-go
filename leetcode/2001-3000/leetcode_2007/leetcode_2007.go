package leetcode_2007

import "sort"

// https://leetcode.cn/problems/find-original-array-from-doubled-array/
func findOriginalArray(changed []int) []int {
	// 1.数组长度是偶数
	if len(changed) == 0 || len(changed)&1 == 1 {
		return nil
	}
	return mapSolution(changed)
}

// 1. map + 排序的解决方案.
// 1. 二倍数组, 优先匹配小值,
func mapSolution(changed []int) []int {

	// 2. check
	var (
		cnts   = make(map[int]int) // 记录数字出现的频次
		result = make([]int, 0)
	)
	for _, num := range changed {
		cnts[num]++
	}
	sort.Ints(changed)
	for _, num := range changed {
		if cnts[num] == 0 {
			continue
		}
		if cnts[num<<1] == 0 { // 不存在num 的2倍数, 不是二倍数组
			return nil
		}
		cnts[num]--
		cnts[num<<1]--
		result = append(result, num)
	}
	return result
}

// 2. queue + sort solution
// 与上面的一样, 优先匹配小值, 所以可以使用排序后的数组, 利用 队列的方式, 优先匹配小值.
func queueSolution(changed []int) []int {

	// 2. check
	var (
		queue  = make([]int, 0)
		result = make([]int, 0)
	)
	sort.Ints(changed)
	for _, num := range changed {
		//if len(queue) > 0 && num == queue[0]<<1 {
		//	// 遇到2倍数字, 找到一组匹配的值,出队
		//	result = append(result, queue[0])
		//	queue = queue[1:]
		//} else {
		//	queue = append(queue, num)
		//}
		if len(queue) == 0 {
			queue = append(queue, num)
		} else {
			if num == queue[0]<<1 {
				// 遇到2倍数字, 找到一组匹配的值,出队
				result = append(result, queue[0])
				queue = queue[1:]
			} else {
				// 入队
				queue = append(queue, num)
			}
		}
	}
	if len(queue) > 0 {
		return nil
	}
	return result
}
