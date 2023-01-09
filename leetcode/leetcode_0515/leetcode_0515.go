package leetcode_0515

import (
	. "github.com/majunmin/leetcode-go/common"
	"math"
)

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/#/description
func largestValues(root *TreeNode) []int {
	// paramcheck
	if root == nil {
		return nil
	}
	return dsfSolution(root)
}

func dsfSolution(root *TreeNode) []int {
	var result []int
	dfs(root, &result, 0)
	return result
}

func dfs(root *TreeNode, result *[]int, level int) {
	// terminate 前序遍历
	if len(*result) <= level {
		*result = append(*result, root.Val)
	} else {
		if root.Val > (*result)[level] {
			(*result)[level] = root.Val
		}
	}
	if root.Left != nil {
		dfs(root.Left, result, level+1)
	}
	if root.Right != nil {
		dfs(root.Right, result, level+1)
	}
}

func bfsSolution(root *TreeNode) []int {
	// param check
	//if root == nil {
	//	return nil
	//}
	var result []int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		maxVal := math.MinInt
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxVal {
				maxVal = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, maxVal)
	}
	return result
}
