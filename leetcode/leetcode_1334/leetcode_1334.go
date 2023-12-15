package leetcode_1334

import "math"

// https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	return solution1(n, edges, distanceThreshold)
}

// 方法 一: 朴素  Dijstra 算法, 使用链接矩阵存储图
func solution1(n int, edges [][]int, distanceThreshold int) int {
	// init graph
	// 原始 graph
	mp := make([][]int, n)
	// 各个点之间的最短距离
	dis := make([][]int, n)
	// 记录节点是否被访问过
	vis := make([][]bool, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, n)
		dis[i] = make([]int, n)
		vis[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			mp[i][j] = math.MaxInt >> 1
			dis[i][j] = math.MaxInt >> 1
		}
	}
	for _, e := range edges {
		from, to, weight := e[0], e[1], e[2]
		mp[from][to], mp[to][from] = weight, weight
	}

	// process
	for i := 0; i < n; i++ {
		vis[i][i] = true
		dis[i][i] = 0
		for cnt := 1; cnt < n; cnt++ { // n个点，需要循环n-1次, 一次确定一个最近的点
			t := -1
			for k := 0; k < n; k++ {
				if vis[i][k] {
					continue
				}
				if t == -1 || dis[i][k] < dis[i][t] {
					t = k
				}
			}

			// [i-t] 的最短距离 是 dis[i][t]
			for k := 0; k < n; k++ {
				dis[i][k] = min(dis[i][k], dis[i][t]+mp[t][k])
			}
			vis[i][t] = true
		}
	}

	// ans[0]: 距离  ans[1]:节点数
	ans := []int{math.MaxInt, 0}
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if dis[i][j] <= distanceThreshold {
				cnt++
			}
		}

		if cnt <= ans[0] {
			ans[0], ans[1] = cnt, i
		}
	}
	return ans[1]
}
