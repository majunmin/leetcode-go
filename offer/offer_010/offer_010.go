package offer_010

//https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/?envType=study-plan&id=lcof
func fib(n int) int {
	//// terminate
	//if n == 0 || n == 1 {
	//	return n
	//}
	//// process
	//return fib(n-1) + fib(n-2)

	return flagSolution(n)
}

func flagSolution(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	p1, p2 := 0, 1
	res := 1
	for i := 2; i <= n; i++ {
		res = p1 + p2
		p1, p2 = p2, res
	}
	return res % (1e9 + 7)
}

func arraySolution(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	nums := make([]int, n+1)
	nums[0], nums[1] = 0, 1
	for i := 2; i <= n; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums[n]
}
