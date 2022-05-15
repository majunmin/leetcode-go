package leetcode_0022

import "fmt"

//https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	resMap := map[int][]string{
		0: {""},
		1: {"()"},
	}
	for i := 2; i <= n; i++ {
		curRes := make([]string, 0)
		for j := 0; j < i; j++ {
			fmt.Println(j)
			for _, v1 := range resMap[j] {
				for _, v2 := range resMap[i-j-1] {
					curRes = append(curRes, "("+v1+")"+v2)
				}
			}
		}
		resMap[i] = curRes
	}
	return resMap[n]
}

func recurSolution(n int) []string {
	var result []string
	var path []byte
	generateParentHelper(n, 0, 0, &path, &result)
	return result
}
func generateParentHelper(n int, left int, right int, path *[]byte, result *[]string) {
	// terminate
	if left+right == 2*n {
		*result = append(*result, string(*path))
		return
	}

	// Process current
	if left < n {
		*path = append(*path, '(')
		generateParentHelper(n, left+1, right, path, result)
		*path = (*path)[:len(*path)-1]
	}
	if left > right {
		*path = append(*path, ')')
		generateParentHelper(n, left, right+1, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
