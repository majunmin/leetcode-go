package leetcode_0001

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	//param check
	if len(nums) < 2 {
		return nil
	}
	// store num -> idx
	m := make(map[int]int)
	for i, num := range nums {
		if idx, exist := m[target-num]; exist {
			return []int{idx, i}
		}
		m[num] = i
	}
	// 没找到结果
	return nil
}
