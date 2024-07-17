package leetcode_2959

// https://leetcode.cn/problems/number-of-possible-sets-of-closing-branches/description/
func numberOfSets(n int, maxDistance int, roads [][]int) int {
	// 二进制枚举 + Floyd 最短路径算法
	var (
		result int
		d      = make([][]int, n)
	)

	//二进制枚举, mask == 1的位为未关闭的道路, mask == 0的位为关闭的道路
	for mask := 0; mask < 1<<n; mask++ {
		// init d
		for i := 0; i < len(d); i++ {
			d[i] = make([]int, n)
			for j := 0; j < n; j++ {
				d[i][j] = 10000
			}
		}
		for _, road := range roads {
			src, dst, cost := road[0], road[1], road[2]
			d[src][dst] = cost
			d[dst][src] = cost
		}
		// Floyd-Warshall algorithm
		for k := 0; k < n; k++ {
			if (1<<k)&mask == 0 {
				continue
			}
			for i := 0; i < n; i++ {
				if (1<<i)&mask == 0 {
					continue
				}
				for j := i + 1; j < n; j++ {
					if (1<<j)&mask == 0 {
						continue
					}
					d[i][j] = min(d[i][j], d[i][k]+d[k][j])
					d[j][i] = min(d[i][j], d[i][k]+d[k][j])
				}
			}
		}

		// validate
		if isValid(maxDistance, mask, d) {
			result++
		}
	}
	return result
}

func isValid(maxDistance int, mask int, d [][]int) bool {
	for i := 0; i < len(d); i++ {
		if (1<<i)&mask == 0 {
			continue
		}
		for j := i + 1; j < len(d); j++ {
			if (1<<j)&mask == 0 {
				continue
			}
			if d[i][j] > maxDistance {
				return false
			}
		}
	}

	return true
}
