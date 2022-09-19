package offer_035

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodesMap := make(map[*Node]*Node)

	dummyNode := &Node{}
	newHead := dummyNode
	node := head
	for node != nil {
		newHead.Next = &Node{
			Val: node.Val,
		}
		newHead = newHead.Next

		nodesMap[node] = newHead
		node = node.Next
	}

	node = head
	for head != nil {
		nodesMap[head].Random = nodesMap[head.Random]
		head = head.Next
	}

	return dummyNode.Next

}
