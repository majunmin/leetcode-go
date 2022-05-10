package leetcode_0094

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/binary-tree-inorder-traversal/
func inorderTraversal(root *TreeNode) []int {
	//var result []int
	//inorder(root, &result)
	//return result
	return stackSolution(root)
}

func stackSolution(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			// 入栈
			stack = append(stack, node)
			node = node.Left
		}

		// 出栈
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

// recursiveSolution
func inorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, result)
	*result = append(*result, root.Val)
	inorder(root.Right, result)
}
