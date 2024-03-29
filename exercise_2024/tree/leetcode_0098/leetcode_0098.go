package leetcode_0098

import (
	"math"

	. "github.com/majunmin/leetcode-go/common"
)

func isValidBST(root *TreeNode) bool {
	// param check
	if root == nil {
		return true
	}
	return isValidBSTHelper(root, math.MinInt, math.MaxInt)
}

func isValidBSTHelper(node *TreeNode, minVal int, maxVal int) bool {
	if node == nil {
		return true
	}
	if node.Val <= minVal || node.Val >= maxVal {
		return false
	}
	return isValidBSTHelper(node.Left, minVal, node.Val) && isValidBSTHelper(node.Right, node.Val, maxVal)
}
