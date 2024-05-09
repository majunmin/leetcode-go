package leetcode_1863

import "fmt"

// https://leetcode.cn/problems/sum-of-all-subset-xor-totals/
func subsetXORSum(nums []int) int {
	//通过位运算计算出所有子集 - 查看 78题
	var (
		size   = 1 << len(nums) // 子集的个数
		result int
	)

	for i := 0; i < size; i++ {
		cur := 0 // 当前子集的异或和
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				cur ^= nums[j]
			}
		}
		result += cur
	}
	return result
}

func backtraceSolution(nums []int) int {
	var (
		result     int
		subsetsXor = make([]int, 0)
	)
	backTrace(nums, -1, 0, &subsetsXor)
	fmt.Println(subsetsXor)
	for i := 0; i < len(subsetsXor); i++ {
		result += subsetsXor[i]
	}
	return result
}

func backTrace(nums []int, begin int, path int, result *[]int) {
	*result = append(*result, path)
	for i := begin; i < len(nums); i++ {
		backTrace(nums, i+1, path^nums[i], result)
	}
}
