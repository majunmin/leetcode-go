package offer_053

//https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/?envType=study-plan&id=lcof
func missingNumber(nums []int) int {

	// 二分查找
	l := len(nums)
	if l == 0 {
		return -1
	}
	left, right := 0, l-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == mid {
			left = mid + 1
		} else {
			right = mid
		}
	}
	//return left
	if left != nums[left] {
		return left
	}
	return left + 1
}
