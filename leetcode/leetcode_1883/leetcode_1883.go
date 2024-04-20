package leetcode_1883

import "math"

func minSkips(dist []int, speed int, hoursBefore int) int {

	var (
		n     = len(dist)
		dists = append([]int{0}, dist...)
		dp    = make([][]float64, n+1)
	)
	// 状态定义
	// dp[i][j] 第i条路 跳过j次, 的耗时
	// dp[i][j] = min(skip, not skip)
	// dp[i][j] = min(dp[i-1][j-1] + dist[i]*1.0 / speed, dp[i-1][j] + ceil(dists[i]*1.0/ speed))

	//for i := 0; i <= n; i++ {
	//	total += dists[i]
	//}
	//if float64(total)/float64(speed) > float64(hoursBefore) {
	//	return -1
	//}

	// 状态初始化
	for i := 0; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = 1e20
		}
	}
	dp[0][0] = 0
	for i := 1; i < n; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = math.Ceil(dp[i-1][j] + float64(dists[i])/float64(speed))
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+float64(dists[i]*1.0)/float64(speed))
			}
		}
	}

	for i := 0; i <= n; i++ {
		dp[n][i] = dp[n-1][i] + float64(dists[n])/float64(speed)
		if dp[n][i] <= float64(hoursBefore) {
			return i
		}
	}
	return -1
}

// 如何避免除法运算? 避免浮点数误差.
// 将道路的长度 和 hoursBefore 长度均乘以 speed.
//
//	时间 x 的下一个 speed 的倍数小时为: (x - 1 )/ speed + 1
func minSkips2(dist []int, speed int, hoursBefore int) int {
	var (
		n  = len(dist)
		dp = make([][]int, n+1)
	)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = 1 << 32
		}
	}
	dp[0][0] = 0
	// 递推方程
	for i := 1; i <= n; i++ {
		for j := 0; j <= i; j++ {
			// 不跳过
			if i != j {
				dp[i][j] = min(dp[i][j], ((dp[i-1][j]+dist[i-1]*speed-1)/speed+1)*speed)
			}

			if j > 0 {
				// 跳过
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+dist[i-1])
			}
		}
	}
	for i := 0; i <= n; i++ {
		if dp[n-1][i] <= hoursBefore*speed {
			return i
		}
	}
	return -1
}
