package leetcode_0092

import . "github.com/majunmin/leetcode-go/common"

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// param check
	if left > right {
		panic("invalid left > right")
	}
	if left == right {
		return head
	}
	dummyNode := &ListNode{}
	dummyNode.Next = head
	pre := dummyNode
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	// 找到 end 这段代码 是判断 right是否超出阈值的.
	// 也可以去掉, 在 reverse 中 判断  prev.Next != nil && prev.Next.Next != nil
	end := pre
	for i := left; i <= right; i++ {
		if end == nil {
			panic(" right is out of range.")
		}
		end = end.Next
	}

	// reverse
	for i := left + 1; i <= right; i++ {
		cur := pre.Next
		tmp := cur.Next
		cur.Next = cur.Next.Next
		pre.Next = tmp
		tmp.Next = cur
	}
	return dummyNode.Next
}
