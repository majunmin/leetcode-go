package leetcode_0002

import . "github.com/majunmin/leetcode-go/common"

// https: //leetcode.cn/problems/add-two-numbers/?envType=study-plan-v2&envId=top-100-liked
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var dummyNode ListNode
	prev := &dummyNode
	p := 0
	for l1 != nil || l2 != nil || p > 0 {
		curVal := p
		if l1 != nil {
			curVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curVal += l2.Val
			l2 = l2.Next
		}
		val := curVal % 10
		p = curVal / 10
		// 拼接 result list
		prev.Next = &ListNode{Val: val}
		prev = prev.Next
	}
	//if p == 1 {
	//	prev.Next = &ListNode{Val: 1}
	//}
	return dummyNode.Next
}
