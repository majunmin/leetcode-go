package offer_028

import . "github.com/majunmin/leetcode-go/offer"

// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/?envType=study-plan&id=lcof
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return symmetric(root.Left, root.Right)

}

func symmetric(node1, node2 *TreeNode) bool {
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	}

	if node1 == nil || node2 != nil {
		return true
	}

	if node1.Val != node2.Val {
		return false
	}

	return symmetric(node1.Left, node2.Right) && symmetric(node1.Right, node2.Left)

}
