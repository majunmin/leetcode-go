// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 21:17
package leetcode_0110

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode-cn.com/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	_, isBalance := isBalancedTree(root)
	return isBalance
}

func isBalancedTree(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}
	if node.Left == nil && node.Right == nil {
		return 1, true
	}
	lHeight, b1 := isBalancedTree(node.Left)
	rHeight, b2 := isBalancedTree(node.Right)
	return maxInt(lHeight, rHeight), b1 && b2 && heightDiff(lHeight, rHeight) <= 1

}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func heightDiff(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}
