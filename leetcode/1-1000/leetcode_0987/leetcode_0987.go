package leetcode_0987

import (
	"math"
	"sort"

	. "github.com/majunmin/leetcode-go/common"
)

type item struct {
	rowIdx int
	colIdx int
	val    int
}

func verticalTraversal(root *TreeNode) [][]int {
	// param check
	if root == nil {
		return nil
	}

	items := make([]*item, 0)

	dfs(root, &items, 0, 0)
	sort.Slice(items, func(i, j int) bool {
		if items[i].colIdx == items[j].colIdx {
			return items[i].rowIdx < items[j].rowIdx
		}
		return items[i].colIdx < items[j].colIdx
	})

	var result [][]int
	lastCol := math.MinInt32
	for _, node := range items {
		if node.colIdx != lastCol {
			lastCol = node.colIdx
			result = append(result, make([]int, 0))
		}
		result[len(result)-1] = append(result[len(result)-1], node.val)
	}
	return result
}

func dfs(node *TreeNode, items *[]*item, rowIdx, colIdx int) {
	if node == nil {
		return
	}
	*items = append(*items, &item{
		rowIdx: rowIdx,
		colIdx: colIdx,
		val:    node.Val,
	})
	dfs(node.Left, items, rowIdx+1, colIdx-1)
	dfs(node.Right, items, rowIdx+1, colIdx+1)
}
