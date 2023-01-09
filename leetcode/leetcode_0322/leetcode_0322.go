package leetcode_0322

import (
	"math"
)

// https://leetcode.cn/problems/coin-change/description/
func coinChange(coins []int, amount int) int {
	//return recurSolution(coins, amount, make([]int, amount+1))
	return dpSolution(coins, amount)
}

// dpSolution
// 思路比较好理解， 类似于爬楼梯 问题
// - 理解无后效性
func dpSolution(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = minInt(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// 如果没有加上 记忆化,重复计算, 超时
func recurSolution(coins []int, amount int, cntArr []int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	//if len(coins) == 0 {
	//	return -1
	//}
	if cntArr[amount] != 0 {
		return cntArr[amount]
	}
	minCnt := math.MaxInt32
	for _, coin := range coins {
		cnt := recurSolution(coins, amount-coin, cntArr)
		if cnt >= 0 && cnt < minCnt {
			minCnt = cnt
		}
	}
	if minCnt == math.MaxInt32 {
		cntArr[amount] = -1
	}
	cntArr[amount] = minCnt + 1
	return cntArr[amount]
}

// bfsSolution   OOM
func bfsSolution(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	level := 1
	queue := coins
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curCoin := queue[i]
			if curCoin == amount {
				return level
			}
			if curCoin > amount {
				continue
			}
			for _, coin := range coins {
				queue = append(queue, curCoin+coin)
			}
		}
		queue = queue[size:]
		level++
	}
	return -1
}
