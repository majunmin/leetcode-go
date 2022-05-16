// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-17 22:31
package leetcode_0070

//https://leetcode-cn.com/problems/climbing-stairs/
func climbStairs(n int) int {
	return iterSolution2(n)
}

// 利用数组记忆前面的计算, 避免重复计算
func iterSolution2(n int) int {
	arr := make([]int, n+1)
	arr[0], arr[1] = 1, 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

// 利用双指针优化空间占用
func iterSolution1(n int) int {
	//双指针法, 当前结果 依赖于 前两个结果
	//p, q, result := 1, 1, 1
	//for i := 2; i <= n; i++ {
	//	result = p + q
	//	p = q
	//	q = result
	//}
	//return result
	p, q := 1, 1
	for i := 2; i <= n; i++ {
		p = p + q
		// swap p, q
		p, q = q, p
	}

	return q
}

func recursionSolution(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
