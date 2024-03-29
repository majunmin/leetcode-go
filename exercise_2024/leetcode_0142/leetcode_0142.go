package leetcode_0142

import . "github.com/majunmin/leetcode-go/common"

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	// process
	var flag bool
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			flag = true
			break
		}
	}
	// 没有交点
	if !flag {
		return nil
	}
	// find 交点
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
