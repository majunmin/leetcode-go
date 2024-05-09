package leetcode_1310

// https://leetcode.cn/problems/xor-queries-of-a-subarray/
func xorQueries(arr []int, queries [][]int) []int {
	// param check
	if len(arr) == 0 || len(queries) == 0 {
		return nil
	}
	// 类似于前缀和的解法
	preX := make([]int, len(arr)+1)
	for i := range arr {
		preX[i+1] = preX[i] ^ arr[i]
	}
	result := make([]int, 0, len(queries))
	for _, item := range queries {
		l, r := item[0], item[1]
		// (l, r]
		result = append(result, preX[l]^preX[r+1])
	}
	return result
}

func lowbit(x int) int {
	return x & -x
}

func update(arr []int, n, val int) {
	for i := n; i < len(arr); i += lowbit(i) {
		arr[i] ^= val
	}
}

func q(arr []int, n int) int {
	var result int
	for i := n; i > 0; i -= lowbit(i) {
		result ^= arr[i]
	}
	return result
}

func query(arr []int, l, r int) int {
	return q(arr, r+1) ^ q(arr, l)
}

// 树状数组的解法
func xorQueries2(arr []int, queries [][]int) []int {
	// param check
	if len(arr) == 0 || len(queries) == 0 {
		return nil
	}
	tree := make([]int, len(arr)+1)
	for i, num := range arr {
		update(tree, i+1, num)
	}
	result := make([]int, 0, len(queries))
	for _, item := range queries {
		result = append(result, query(tree, item[0], item[1]))
	}
	return result
}
