package exercise_2024

import . "github.com/majunmin/leetcode-go/common"

func inorderTraversal(root *TreeNode) []int {
	// param check
	if root == nil {
		return nil
	}
	var (
		node   = root
		result []int

		stack = make([]*TreeNode, 0)
	)
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// pop stack
		node = stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		node = nil
	}
	return result
}
