package leetcode_0144

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/binary-tree-preorder-traversal/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var result []int
	preorder(root, &result)
	return result
}

// recursive solution
func preorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorder(root.Left, result)
	preorder(root.Right, result)
}

// stack solution
func preorderStack(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			result = append(result, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		// 出栈
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return result
}
