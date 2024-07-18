package leetcode_0427

//https://leetcode.cn/problems/construct-quad-tree/?envType=study-plan-v2&envId=top-interview-150
/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	//param check
	n := len(grid)
	if n == 0 {
		return nil
	}
	return build(grid, 0, n-1, 0, n-1)
}

func build(grid [][]int, l int, r int, top int, down int) *Node {
	// check
	same := true
	for i := l; i <= r; i++ {
		for j := top; j <= down; j++ {
			if grid[i][j] != grid[l][top] {
				same = false
			}
		}
		if !same {
			break
		}
	}
	node := new(Node)
	if !same {
		node.IsLeaf = true
		node.Val = grid[l][top] == 1
		return node
	}
	// terminate
	node.IsLeaf = false
	mid1, mid2 := l+(r-l)>>1, top+(down-top)>>1
	node.TopLeft = build(grid, l, mid1, top, mid2)
	node.TopRight = build(grid, mid1+1, r, top, mid2)
	node.BottomLeft = build(grid, l, mid1, mid2+1, down)
	node.BottomRight = build(grid, mid1+1, r, mid2+1, down)
	node.Val = true //(true | false)都可
	return node
}
