package leetcode_2386

import (
	"container/heap"
	"slices"
	"sort"
)

// pq item [2]int{num, idx}
type pq [][2]int64

func newPQ() *pq {
	var p pq
	heap.Init(&p)
	return &p
}

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x any) {
	*p = append(*p, x.([2]int64))
}

func (p *pq) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func (p *pq) push(x any) {
	heap.Push(p, x)
}

func (p *pq) pop() any {
	return heap.Pop(p)
}

// https://leetcode.cn/problems/find-the-k-sum-of-an-array/
func kSum(nums []int, k int) int64 {
	// param check
	var maxSum int64
	for i, num := range nums {
		if num >= 0 {
			maxSum += int64(num)
		} else {
			nums[i] = -nums[i]
		}
	}
	if k == 1 {
		return maxSum
	}
	sort.Ints(nums)
	pQueue := newPQ()
	pQueue.push([2]int64{int64(nums[0]), 0})
	// 转化为 求 nums 的第 k-1 小的子序列和
	for i := 0; i < k-1; i++ {
		top := pQueue.pop().([2]int64)
		sum, prevI := top[0], top[1]
		if i == k-2 {
			return maxSum - sum
		}
		if int(prevI) < len(nums)-1 {
			pQueue.push([2]int64{sum + int64(nums[prevI+1]), int64(prevI) + 1})
			pQueue.push([2]int64{sum - int64(nums[prevI]) + int64(nums[prevI+1]), int64(prevI) + 1})
		}
	}
	return -1
}

// 由于k 的取值范围 ,时间复杂度 为 2^n, 超时
func backtraceSolution(nums []int, k int) int64 {
	// param check
	if k > 1<<len(nums) {
		panic("invalid param")
	}
	var sums []int64
	sort.Ints(nums)
	backtrace(&sums, nums, 0, []int64{})
	slices.SortFunc(sums, func(a, b int64) int {
		return int(a - b)
	})
	return sums[k-1]
}

func backtrace(sums *[]int64, nums []int, begin int, path []int64) {
	// add to result
	var sum int64
	for _, item := range path {
		sum += item
	}
	*sums = append(*sums, sum)

	if begin == len(nums) {
		return
	}

	for i := begin; i < len(nums); i++ {
		// 同层不能重复
		if i > begin && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, int64(nums[i]))
		backtrace(sums, nums, i+1, path)
		path = path[:len(path)-1]
	}
}
