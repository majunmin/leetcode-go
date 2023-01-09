package leetcode_0496

import "math"

// https://leetcode.cn/problems/next-greater-element-i/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nil
	}

	m := make(map[int]int)

	stack := make([]int, 0)
	newNums2 := make([]int, 0, len(nums2)+1)
	copy(newNums2[1:], nums2)
	newNums2[0] = math.MaxInt
	for i := 1; i < len(newNums2); i++ {
		for len(stack) > 0 && newNums2[i] > stack[len(stack)-1] {
			// pop
			m[stack[len(stack)-1]] = newNums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, newNums2[i])
	}
	res := make([]int, 0, len(nums1))
	for _, n := range nums1 {
		if v, exist := m[n]; exist {
			res = append(res, v)
		} else {
			res = append(res, -1)
		}

	}
	return res
}
