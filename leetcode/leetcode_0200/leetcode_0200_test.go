package leetcode_0200

import (
	"fmt"
	"testing"
)

func TestNumOfIsland(t *testing.T) {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '1'},
	}))
}

func TestUnifind(t *testing.T) {
	uf := NewUnionFind(10)
	fmt.Println(uf.Count())
	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(5, 6)
	fmt.Println(uf.Count())
	uf.Union(1, 3)
	fmt.Println(uf.Count())
}
