package leetcode_100240

import (
	"math"

	"github.com/erriapo/redblacktree"
)

func minimumDistance(points [][]int) int {
	// 曼哈顿距离 转化为 切比雪夫距离
	xs := redblacktree.NewTree()
	ys := redblacktree.NewTree()
	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		xs.Put(x, x)
		ys.Put(y, y)
	}
	var (
		result = math.MaxInt
	)

	for _, p := range points {
		x, y := p[0]+p[1], p[1]-p[0]
		xs.Delete(x)
		ys.Delete(y)
		result = min(result, max())
		xs.Put(x, x)
		ys.Put(y, y)
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
