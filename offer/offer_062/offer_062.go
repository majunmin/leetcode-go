package offer_062

// 此题目可以利用数学解法
func lastRemaining(n int, m int) int {
	return f(n, m)
}

func f(n int, m int) int {
	if n == 1 {
		return 0
	}
	//
	return (f(n-1, m) + m) % n
}

// 执行超时
func arrSolution(n int, m int) int {
	// param check
	if n < 1 {
		panic("n must be greater than 0")
	}
	if n == 1 {
		return 0
	}
	// 弄成一个循环数组.
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	//
	preIdx := -1
	for len(arr) > 1 {
		delIdx := (preIdx + m) % len(arr)
		arr = append(arr[:delIdx], arr[delIdx+1:]...)
		preIdx = delIdx - 1
	}
	return arr[0]
}
