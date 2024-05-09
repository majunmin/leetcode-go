package leetcode_1457

import (
	"fmt"

	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/pseudo-palindromic-paths-in-a-binary-tree/
func pseudoPalindromicPaths(root *TreeNode) int {
	counter := make([]int, 10)
	return preOrder(root, counter)
}

func preOrder(root *TreeNode, counter []int) int {
	if root == nil {
		return 0
	}
	var res int
	counter[root.Val]++
	if root.Left == nil && root.Right == nil {
		if isPalindromic(counter) {
			res = 1
		}
	} else {
		res = preOrder(root.Left, counter) + preOrder(root.Right, counter)
	}
	counter[root.Val]--
	return res
}

func isPalindromic(counter []int) bool {
	fmt.Println(counter)
	singleNum := 0
	for _, item := range counter {
		if item&1 == 1 {
			singleNum++
		}
	}
	return singleNum <= 1
}
