package leetcode_0001

// https://leetcode.cn/problems/two-sum/?envType=study-plan-v2&envId=top-100-liked
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums)) // num -> idx

	for i, num := range nums {
		if idx, exist := numsMap[target-num]; exist {
			return []int{idx, i}
		}
		numsMap[num] = i
	}
	// not found
	return nil
}
