package __doublepointer

// 快慢指针解法
// 类似环形链表, 找到其相遇的点
func findDuplicate(nums []int) int {
	var (
		slow = nums[0]
		fast = nums[nums[0]]
	)
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 找出相遇点
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
