package leetcode_1521

import (
	"maps"
	"math"
)

// https://leetcode.cn/problems/find-a-value-of-a-mysterious-function-closest-to-target/
func closestToTarget(arr []int, target int) int {
	// return min diff
	resultSet := make(map[int]bool)
	//生成所有结果
	for i := 0; i < len(arr); i++ {
		curLevel := make(map[int]bool)
		maps.Copy(curLevel, resultSet)
		num := arr[i]
		for num2 := range curLevel {
			resultSet[num&num2] = true
		}
		resultSet[num] = true
	}
	// 寻找结果
	result := math.MaxInt
	for num := range resultSet {
		result = min(result, abs(num-target))
	}
	return result
}

func abs(num int) int {
	if num < 0 {
		return -num

	}
	return num
}
