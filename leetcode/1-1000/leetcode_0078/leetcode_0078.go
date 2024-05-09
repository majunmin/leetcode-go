// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-24 22:39
package leetcode_0078

// subsets
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	backtrace(0, []int{}, nums, &result)
	return result
}

func backtrace(begin int, path []int, nums []int, result *[][]int) {
	dst := make([]int, len(path))
	copy(dst, path)
	*result = append(*result, dst)

	for i := begin; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrace(i+1, path, nums, result)
		path = path[:len(path)-1]
	}
}

// 位运算求解
func bitSolution(nums []int) [][]int {
	// 位运算
	// 0-8 [000, 001, 010,011,100,101,110,111]
	result := make([][]int, 0)
	for i := 0; i < pow(2, len(nums)); i++ {
		i := i
		path := make([]int, 0, len(nums))
		for j := 0; j < len(nums); j++ {
			if i&1 == 1 {
				path = append(path, nums[j])
			}
			i = i >> 1
		}

		result = append(result, path)
	}
	return result
}

func pow(num, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return num
	}
	if m&1 == 1 {
		return pow(num, m/2) * pow(num, m/2+1)
	}
	return pow(num, m/2) * pow(num, m/2)
}
