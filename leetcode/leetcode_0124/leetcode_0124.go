package leetcode_0124

import (
	"math"

	. "github.com/majunmin/leetcode-go/common"
)

func maxPathSum(root *TreeNode) int {
	var maxSum int
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return math.MinInt8
		}
		// 只有 gain  > 0 时(有增益) 才会被计入
		leftGain, rightGain := max(maxGain(node.Left), 0), max(maxGain(node.Right), 0)
		// 更新结果
		maxSum = max(maxSum, leftGain+rightGain+node.Val)
		// 返回 路径最大增益
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
