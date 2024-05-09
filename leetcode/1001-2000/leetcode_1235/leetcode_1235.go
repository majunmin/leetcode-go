package leetcode_1235

import "slices"

// oom, 需要优化空间
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	// dp[i] = max(dp[i-1], dp[e2s[i]])
	// param check
	var (
		n    = len(startTime)
		jobs = make([][3]int, n)
	)
	for i := 0; i < n; i++ {
		jobs = append(jobs, [3]int{startTime[i], endTime[i], profit[i]})
	}
	slices.SortFunc(jobs, func(a, b [3]int) int {
		return a[1] - b[1]
	})
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		k := binarySearch(jobs, i-1, jobs[i-1][0])
		dp[i] = max(dp[i-1], dp[k]+jobs[i-1][2])
	}
	return dp[n]
}

// 找到 endTIme <upper 的最大下标
func binarySearch(jobs [][3]int, right int, target int) int {
	var left int
	for left < right {
		mid := left + (right-left)>>1
		if jobs[mid][1] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func solution1(startTime []int, endTime []int, profit []int) int {
	// 动态规划, 维护一个  e2s = endTime -> {startTime, profile} 的map
	// dp[i] = max(dp[i-1], dp[e2s[i]])
	// param check
	if len(startTime) == 0 {
		return 0
	}
	var (
		N   = 0
		dp  = make([]int, N+1)
		e2s = make(map[int]map[int]int) // endTime -> {startTime, profit}
	)
	// 1. build e2s
	for i := 0; i < len(startTime); i++ {
		N = max(N, endTime[i])
		if _, exist := e2s[endTime[i]]; !exist {
			e2s[endTime[i]] = make(map[int]int)
			e2s[endTime[i]][startTime[i]] = profit[i]
		} else if _, exist = e2s[endTime[i]][startTime[i]]; !exist {
			e2s[endTime[i]][startTime[i]] = profit[i]
		} else {
			e2s[endTime[i]][startTime[i]] = max(e2s[endTime[i]][startTime[i]], profit[i])
		}
	}
	// init state
	//dp[1] == dp[0] = 0
	for i := 2; i <= N; i++ {
		if _, exist := e2s[i]; !exist {
			dp[i] = dp[i-1]
		}
		dp[i] = dp[i-1]
		// startTime profit
		for s, p := range e2s[i] {
			dp[i] = max(dp[i], dp[s]+p)
		}
	}
	return dp[len(dp)-1]
}
