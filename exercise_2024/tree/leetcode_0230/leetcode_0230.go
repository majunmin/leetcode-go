package leetcode_0230

import . "github.com/majunmin/leetcode-go/common"

func kthSmallest(root *TreeNode, k int) int {
	var result int
	var search func(node *TreeNode)
	search = func(node *TreeNode) {
		search(node.Left)
		k--
		if k == 0 {
			result = node.Val
			return
		}
		search(node.Right)
	}
	search(root)
	return result
}

func inorderSolution(root *TreeNode, k int) int {
	// 二叉树的中序遍历, 是有序的， 遍历到的 第k个节点
	var (
		node  = root
		stack = make([]*TreeNode, 0)
	)
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// pop
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			break
		}
		if node.Right != nil {
			node = node.Right
		} else {
			node = nil
		}
	}
	if k == 0 {
		return node.Val
	}
	panic("invalid param")
}
