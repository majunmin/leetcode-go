package offer_055

import . "github.com/majunmin/leetcode-go/offer"

// https://leetcode.cn/problems/ping-heng-er-cha-shu-lcof/?envType=study-plan&id=lcof
// 左子树 是 平衡二叉树  && 右子树是平衡二叉树  && 左右子树 高度相差 < 1
func isBalanced(root *TreeNode) bool {
	// param check
	if root == nil {
		return true
	}

	//
	balanced(root)
	leftBalanced, leftDep := balanced(root.Left)
	rightBalanced, rightDep := balanced(root.Right)
	return leftBalanced && rightBalanced && abs(rightDep-leftDep) < 2

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func balanced(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	lBalanced, lDep := balanced(node.Left)
	rBalanced, rDep := balanced(node.Right)

	return lBalanced && rBalanced && abs(lDep-rDep) < 2, maxInt(lDep, rDep) + 1
}
