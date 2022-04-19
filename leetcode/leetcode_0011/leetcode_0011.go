// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-17 21:57
package leetcode_0011

//https://leetcode-cn.com/problems/container-with-most-water/
// 分析:
// 盛水多少 有  宽度(right - left) * 高度(min(height[left], height[right])) 决定
// 长版向内收缩,面积一定变小,
// 短板向内收缩, 面积面积可能增大
// 1. init left, right 双指针
// 2. for (left < right) 循环收窄
//    - update maxArea
//    - 短板向内收缩
// 3.  返回结果
func maxArea(height []int) int {

	// 盛水最多由  两个柱子之间最短的柱子决定,宽度
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		if height[left] < height[right] {
			maxArea = maxInt(maxArea, (right-left)*height[left])
			left++
			for left < right && height[left] <= height[left-1] {
				left++
			}
		} else {
			maxArea = maxInt(maxArea, (right-left)*height[right])
			right--
			for left < right && height[right] <= height[right+1] {
				right--
			}
		}
	}
	return maxArea
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
