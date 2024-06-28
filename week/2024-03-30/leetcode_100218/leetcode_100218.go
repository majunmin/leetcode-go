package leetcode_100218

import (
	"fmt"
	"math"
	"sort"
)

func sumOfPowers(nums []int, k int) int {
	// param check
	if len(nums) < k {
		return 0
	}
	if k == 1 {
		return 0
	}
	sort.Ints(nums)
	var result int
	for i := 0; i <= len(nums)-k; i++ {
		result += backtrace(nums, k-1, []int{nums[i]}, math.MaxInt, i)
	}
	return result
}

func backtrace(nums []int, k int, path []int, engine int, begin int) int {
	if k == 0 {
		fmt.Println(path)
		return engine
	}
	var res int
	for i := begin + 1; i < len(nums)-k+1; i++ {
		e := min(engine, nums[i]-nums[begin])
		path = append(path, nums[i])
		res += backtrace(nums, k-1, path, e, i)
		path = path[:len(path)-1]
	}
	return res % (109 + 7)
}
