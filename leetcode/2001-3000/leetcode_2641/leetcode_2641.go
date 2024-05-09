package leetcode_2641

import (
	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/cousins-in-binary-tree-ii/
func replaceValueInTree(root *TreeNode) *TreeNode {
	// param check
	if root == nil {
		return nil
	}
	// 分两步
	// 1. 线上提， 算出子节点的和，记录在父节点里面.
	// 2. 将 root.Val = 0, 交换  兄弟节点. 然后将 父节点的值 填入子节点里.

	// bfs 处理 计算层级和
	levelTotal := bfs(root)
	//  dfs 后序遍历, 计算  节点的 值
	dfs(root, 0, levelTotal)
	root.Val = 0
	return root
}

func bfs(root *TreeNode) []int {
	levelTotal := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		total := 0
		for i := 0; i < size; i++ {
			item := queue[i]
			total += item.Val
			childSum := 0
			if item.Left != nil {
				queue = append(queue, item.Left)
				childSum += item.Left.Val
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
				childSum += item.Right.Val
			}
		}
		levelTotal = append(levelTotal, total)
		queue = queue[size:]
	}
	return levelTotal
}

func dfs(root *TreeNode, level int, levelTotal []int) {
	if root.Left == nil && root.Right == nil {
		return
	}
	curSum := 0
	if root.Left != nil {
		dfs(root.Left, level+1, levelTotal)
		curSum += root.Left.Val
	}
	if root.Right != nil {
		dfs(root.Right, level+1, levelTotal)
		curSum += root.Right.Val
	}

	curVal := levelTotal[level+1] - curSum
	// 计算子节点的值
	if root.Left != nil {
		root.Left.Val = curVal
	}
	if root.Right != nil {
		root.Right.Val = curVal
	}
}
