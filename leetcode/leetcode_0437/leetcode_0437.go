package leetcode_0437

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/path-sum-iii/?envType=study-plan-v2&envId=top-100-liked
func pathSum(root *TreeNode, targetSum int) int {
	// param check
	if root == nil {
		return 0
	}

	return rootSum(root, targetSum) +
		pathSum(root.Left, targetSum) +
		pathSum(root.Right, targetSum)
}

// 以 root 为根节点, sum = targetSum 的路径数
func rootSum(node *TreeNode, targetSum int) int {
	// terminate
	if node == nil {
		return 0
	}
	//
	var ret int
	if targetSum == node.Val {
		ret += 1 // 后面的和如果0,也是一条路径. 所以这里+1 继续递归
	}
	return ret + rootSum(node.Left, targetSum-node.Val) + rootSum(node.Right, targetSum-node.Val)
}
