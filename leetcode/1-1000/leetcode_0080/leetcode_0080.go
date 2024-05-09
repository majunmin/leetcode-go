package leetcode_0080

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/?envType=study-plan-v2&envId=top-interview-150
func removeDuplicates(nums []int) int {
	var cur int
	var cnt int
	for i := 0; i < len(nums); i++ {
		if nums[i] == cur {
			cnt++
		} else {
			// 遇到新数字, 重新计数
			cur = nums[i]
			cnt = 1
		}
		if cnt == 3 {
			// remove nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			i--
			cnt--
		}
	}

	return len(nums)
}

func removeDuplicates2(nums []int) int {
	// param check
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		// 这里为什么是判断  nums[fast] != nums[slow - 2 ] ?
		// 检查上上个保存的元素与 当前检查元素 nums[fast] 是否相等
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
