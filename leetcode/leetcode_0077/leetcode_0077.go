// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-25 23:10
package leetcode_0077

func combine(n int, k int) [][]int {
	result := make([][]int, 0)
	backTrace([]int{}, 1, k, n, &result)
	return result
}

func backTrace(path []int, begin, k, n int, result *[][]int) {
	// 剪枝
	if len(path) > k {
		return
	}

	if len(path) == k {
		dst := make([]int, k)
		copy(dst, path)
		*result = append(*result, dst)
	}

	// for choice in choiceList
	for i := begin; i <= n; i++ {
		path = append(path, i)
		backTrace(path, i+1, k, n, result)
		path = path[:len(path)-1]
	}
}
