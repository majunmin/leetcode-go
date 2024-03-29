package exercise_2024

import . "github.com/majunmin/leetcode-go/common"

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}

	dummyNode := &ListNode{}
	dummyNode.Next = head
	return reverseKGroupInner(dummyNode, head, k)
}

func reverseKGroupInner(prev *ListNode, cur *ListNode, k int) *ListNode {
	tail := cur
	for i := 1; i < k; i++ {
		tail = tail.Next
		// 不足k个节点, 直接返回
		if tail == nil {
			return cur
		}
	}
	next := tail.Next
	for i := 1; i < k; i++ {
		// 将 cur.Next 插入到 prev.Next 位置
		tmp := cur.Next
		cur.Next = cur.Next.Next
		tmp.Next = prev.Next
		prev.Next = tmp
	}
	cur.Next = reverseKGroupInner(cur, next, k)
	return prev.Next
}
