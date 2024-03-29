package leetcode_0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/?envType=study-plan-v2&envId=top-100-liked
func copyRandomList(head *Node) *Node {
	// originList 1 -> 2 -> 3
	// 1. create copyNode  1 -> 1' -> 2 -> 2' -> 3 -> 3'
	// 2. copy random node
	// 3. 拆除 copyNode
	prev := head
	for prev != nil {
		tmp := &Node{Val: prev.Val}
		tmp.Next = prev.Next
		prev.Next = tmp
		prev = prev.Next.Next
	}
	// 2. copy random node
	prev = head
	for prev != nil {
		if prev.Random != nil {
			prev.Next.Random = prev.Random.Next
		}
		prev = prev.Next.Next
	}
	// 3. 拆出 copyNode
	dummyNode := &Node{}
	cur := dummyNode
	prev = head
	for prev != nil {
		cur.Next = prev.Next
		cur = cur.Next
		prev.Next = prev.Next.Next
		prev = prev.Next
	}
	return dummyNode.Next
}
