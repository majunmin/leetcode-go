package leetcode_0215

import (
	"math/rand/v2"
	"sort"
)

// solution2
// 超时啦...
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	sort.Ints(nums)
	return quickFind(nums, 0, n-1, n-k)
}

func quickFind(nums []int, l int, r int, k int) int {
	pivotIdx := l + rand.IntN(r-l+1) // rand(n)随机生成 [0, n-1)
	nums[r], nums[pivotIdx] = nums[pivotIdx], nums[r]
	var cnt int
	for i := l; l < r; i++ {
		if nums[i] < nums[pivotIdx] {
			nums[i], nums[cnt] = nums[cnt], nums[i]
			cnt++
		}
	}
	nums[pivotIdx], nums[cnt] = nums[cnt], nums[pivotIdx]
	if cnt == k {
		return nums[cnt]
	} else if cnt > k {
		return quickFind(nums, l, cnt-1, k)
	}
	return quickFind(nums, cnt+1, r, k-cnt)
}
