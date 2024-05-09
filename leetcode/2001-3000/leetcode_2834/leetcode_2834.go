package leetcode_2834

// https://leetcode.cn/problems/find-the-minimum-possible-sum-of-a-beautiful-array/
func minimumPossibleSum(n int, target int) int {
	// 数学计算
	//https://leetcode.cn/problems/find-the-minimum-possible-sum-of-a-beautiful-array/solutions/2668273/zhao-chu-mei-li-shu-zu-de-zui-xiao-he-by-20h1/
	m := target / 2
	if m >= n {
		return ((n + 1) * n >> 1) % (1e9 + 7)
	}
	return (m*(m+1)/2 + (target+target+(n-m+1))*(n-m)/2) / (1e9 + 7)
}

// 模拟算法, 会超时
func moniSolution(n int, target int) int {
	var (
		sum        int
		num        = 1 // 要加入美丽数组的数字
		i          int // 加入美丽数组的数字个数
		excludeMap = make(map[int]bool)
	)

	for i < n {
		if excludeMap[num] {
			num++
			continue
		}
		sum += num
		if num < target {
			excludeMap[target-num] = true
		}
		num++
		i++
	}
	return sum
}
