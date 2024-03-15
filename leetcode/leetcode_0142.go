package leetcode

import . "github.com/majunmin/leetcode-go/common"

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
