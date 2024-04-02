package leetcode_2952

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/
// 视频题解:https://www.bilibili.com/video/BV1FW4y1c7B3/
func minimumAddedCoins(coins []int, target int) int {
	// param check
	if target == 0 {
		return 0
	}
	//从小到大排序
	// 从左到右遍历硬币
	// 1. 如果出现了 数字 x , 那么 也就意味着可以组合的数字可以到达 2x-1
	sort.Ints(coins)
	var (
		result int
		end    int
		i      int
	)
	// 思路和代码 都有点不好想😂
	//贪心算法
	for end < target {
		if i == len(coins) || end+1 < coins[i] {
			// 此时需要添加一个数
			result++
			end = end*2 + 1 // == (end+1)*2-1
		} else {
			end += coins[i]
			i++
		}
	}
	return result
}
