// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 22:47
package leetcode_0142

import . "github.com/majunmin/leetcode-go/common"

//https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	// https://assets.leetcode-cn.com/solution-static/142/142_fig1.png
	// 利用 快慢指针,存在环  && 当两指针相遇时:我们得到:
	// slow = a+b,  fast = a + b + n(b+c)
	// 2 slow = fast => 2(a + b) = (a + b) + n(b+c) => a = c + (n - 1)(b+c)
	//  由此得出结论:  当一个 节点从head 触发 和 slow 节点 每次都向后移动, 最终会在  入环的第一个节点相遇

	fast := head
	slow := head
	for fast != nil && fast.Next != nil && slow != nil {
		slow = slow.Next
		fast = fast.Next.Next
		// 存在环
		if fast == slow {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}

func detectCycleByHash(head *ListNode) *ListNode {
	//利用一个 hash,记忆之前 遍历过得  listNode
	nodeMap := make(map[*ListNode]struct{})
	node := head
	for node != nil {
		if _, exist := nodeMap[node]; exist {
			return node
		}
		nodeMap[node] = struct{}{}
		node = node.Next
	}
	return nil
}
