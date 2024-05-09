// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-21 20:34
package leetcode_0104

import "github.com/majunmin/leetcode-go/common"

func maxDepth(root *common.TreeNode) int {
	// terminate
	if root == nil {
		return 0
	}
	return maxInt(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
