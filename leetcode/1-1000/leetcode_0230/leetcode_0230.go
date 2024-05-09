package leetcode_0230

import . "github.com/majunmin/leetcode-go/common"

func kthSmallest(root *TreeNode, k int) int {
	//中序遍历
	if root == nil || k <= 0 {
		return -1
	}
	stack := make([]*TreeNode, 0)
	var cnt int
	stack = append(stack, root)
	var node *TreeNode
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node.Left)
			node = node.Left
		}
		// node = pop stack
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cnt++
		if cnt == k {
			return node.Val
		}
		//
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return -1
}
