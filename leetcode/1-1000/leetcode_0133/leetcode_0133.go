package leetcode_0133

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	var (

		// originNode to cloneNode
		nodeMap = make(map[*Node]*Node)
	)
	// dfs or bfs
	dfs(nodeMap, node)
	return nodeMap[node]
}

func dfs(nodeMap map[*Node]*Node, cur *Node) {
	if nodeMap[cur] != nil {
		return
	}
	nodeMap[cur] = new(Node)
	nodeMap[cur].Val = cur.Val
	for _, next := range cur.Neighbors {
		dfs(nodeMap, next)
	}
	//
	for _, next := range cur.Neighbors {
		neighbors := nodeMap[cur].Neighbors
		neighbors = append(neighbors, nodeMap[next])
	}
}
