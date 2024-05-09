package leetcode_2642

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	graph := Constructor(4, [][]int{
		{0, 2, 5},
		{0, 1, 2},
		{1, 2, 1},
		{3, 0, 3},
	})
	fmt.Println(graph.ShortestPath(3, 2))
	fmt.Println(graph.ShortestPath(0, 3))
	graph.AddEdge([]int{1, 3, 4})
	fmt.Println(graph.ShortestPath(0, 3))
}
