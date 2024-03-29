package leetcode_0101

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/symmetric-tree/?envType=study-plan-v2&envId=top-100-liked
func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val && isSymmetricHelper(node1.Left, node2.Right) && isSymmetricHelper(node1.Right, node2.Left)
}
