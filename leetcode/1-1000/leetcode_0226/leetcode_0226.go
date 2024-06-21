package leetcode_0226

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/invert-binary-tree/description/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 使用队列也可以
	// 遍历树的所有节点, 执行交换操作
	stack := make([]*TreeNode, 0)
	result := root
	stack = append(stack, root)
	for len(stack) > 0 {
		// pop stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// swap
		item.Left, item.Right = item.Right, item.Left
		if item.Left != nil {
			stack = append(stack, item.Left)
		}
		if item.Right != nil {
			stack = append(stack, item.Right)
		}
	}
	return result
}
