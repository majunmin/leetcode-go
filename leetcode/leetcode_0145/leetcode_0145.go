package leetcode_0145

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode.cn/problems/binary-tree-postorder-traversal/
func postorderTraversal(root *TreeNode) []int {
	//var result []int
	//postorder(root, &result)
	//return result

	return stackSolution(root)
}

// 栈解决
func stackSolution(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	// 记录上一次被访问的节点  sentinel
	var lastNode *TreeNode
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == lastNode {
			result = append(result, node.Val)
			lastNode = node
			stack = stack[:len(stack)-1]
			node = nil
		} else {
			node = node.Right
		}
	}
	return result
}

func postorder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postorder(root.Left, result)
	postorder(root.Right, result)
	*result = append(*result, root.Val)
}
