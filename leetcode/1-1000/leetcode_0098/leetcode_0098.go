package leetcode_0098

import (
	"math"

	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/validate-binary-search-tree/
func isValidBST(root *TreeNode) bool {
	return inorderSolution(root)
}

// 利用 二叉搜索树的  中序遍历的特点
func inorderSolution(root *TreeNode) bool {
	var stack []*TreeNode
	preNum := math.MinInt

	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val <= preNum {
			return false
		}
		preNum = node.Val
		node = node.Right
	}

	return true
}

// 由上至下  验证
func recurSolution(root *TreeNode) bool {
	return validBSTHelper(root, math.MinInt, math.MaxInt)
}

func validBSTHelper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return validBSTHelper(root.Left, lower, root.Val) && validBSTHelper(root.Right, root.Val, upper)
}
