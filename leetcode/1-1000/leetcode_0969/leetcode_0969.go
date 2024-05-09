// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-19 22:35
package leetcode_0969

// https://leetcode-cn.com/problems/pancake-sorting/
func pancakeSort(arr []int) []int {
	// 1. 反转操作  可以观察出 尾部元素可以保持不变
	// 2. 每次 将最大值 翻转到最后边,  可以完成  煎饼排序
	//   - 先将最大值翻转到 index 0, 之后再翻转至 最大值位置
	result := make([]int, 0)
	for i := len(arr) - 1; i > 0; i-- {
		maxValIdx := finMaxIdx(arr, i)
		if maxValIdx == i {
			continue
		}
		result = append(result, maxValIdx+1, i+1)
		reverse(arr, 0, maxValIdx)
		reverse(arr, 0, i)
	}
	return result
}

func finMaxIdx(arr []int, limit int) int {
	maxIdx := 0
	for i := 1; i <= limit; i++ {
		if arr[i] > arr[maxIdx] {
			maxIdx = i
		}
	}
	return maxIdx
}

// 翻转数组
func reverse(arr []int, startIdx, endIdx int) {
	for startIdx < endIdx {
		arr[startIdx], arr[endIdx] = arr[endIdx], arr[startIdx]
		startIdx++
		endIdx--
	}
}
