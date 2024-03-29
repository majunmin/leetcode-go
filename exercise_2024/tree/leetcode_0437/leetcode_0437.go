package leetcode_0437

import . "github.com/majunmin/leetcode-go/common"

func rootSum(node *TreeNode, targetSum int) int {
	if node == nil {
		return 0
	}
	var ret int
	if node.Val == targetSum {
		ret += 1
	}
	return ret +
		rootSum(node.Left, targetSum-node.Val) +
		rootSum(node.Right, targetSum-node.Val)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return rootSum(root, targetSum) +
		pathSum(root.Left, targetSum) +
		pathSum(root.Right, targetSum)
}
