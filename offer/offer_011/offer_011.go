package offer_011

//https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/?envType=study-plan&id=lcof
func minArray(numbers []int) int {
	if len(numbers) == 0 {
		return -1
	}
	if len(numbers) == 1 {
		return numbers[0]

	}
	//二分查找
	left, right := 0, len(numbers)-1
	for left < right {
		mid := left + (right-left+1)/2
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] == numbers[right] {
			right--
		} else {
			left = mid + 1
		}
	}

	return numbers[left]
}
