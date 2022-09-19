package offer_053

//https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/?envType=study-plan&id=lcof
func search(nums []int, target int) int {

	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	var count int
	for i := left; i < len(nums); i++ {
		if nums[i] != target {
			break
		}
		count++
	}
	return count
}
