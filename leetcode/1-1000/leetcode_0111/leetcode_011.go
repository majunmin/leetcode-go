package leetcode_0111

import (
	"math"

	. "github.com/majunmin/leetcode-go/common"
)

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//广度优先遍历
	var stack []*TreeNode
	stack = append(stack, root)

	minDep := 0
	for len(stack) == 0 {
		l := len(stack)
		minDep += 1
		for i := 0; i < l; i++ {
			minDep += 1
			// pop
			node := stack[0]
			stack = stack[1:]
			if node.Left == nil && node.Right == nil {
				return minDep
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
	}
	return minDep
}

func recurSolution(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	minDep := math.MaxInt
	if root.Left != nil {
		minDep = minInt(minDep, minDepth(root.Left))
	}
	if root.Right != nil {
		minDep = minInt(minDep, minDepth(root.Right))
	}
	return minDep + 1
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
