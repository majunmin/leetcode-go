package leetcode_3072

import "slices"

func lowbit(num int) int {
	return num & (-num)
}

// 树状数组 结构
type RangeCount struct {
	counts []int
	size   int
	total  int
}

func NewRangeCount(n int) *RangeCount {
	return &RangeCount{
		counts: make([]int, n),
		size:   n,
	}
}
func (rc *RangeCount) Add(num int) {
	for num < rc.size {
		rc.counts[num] += 1
		num += lowbit(num)
	}
	rc.total++
}

func (rc *RangeCount) CountGT(start int) int {
	return rc.total - rc.pre(start)
}

func (rc *RangeCount) pre(end int) int {
	var count int
	for end > 0 {
		count += rc.counts[end]
		end -= lowbit(end)
	}
	return count
}

// https://leetcode.cn/problems/distribute-elements-into-two-arrays-ii/
func resultArray(nums []int) []int {
	var (
		n        = len(nums)
		sortNums = make([]int, n)
		indexMap = make(map[int]int)
	)
	copy(sortNums, nums)
	slices.Sort(sortNums)
	for i, num := range sortNums {
		if _, exist := indexMap[num]; exist {
			continue
		}
		indexMap[num] = i + 1
	}
	arr1, arr2 := make([]int, 0, n), make([]int, 0, n)
	arr1 = append(arr1, nums[0])
	arr2 = append(arr2, nums[1])
	tree1, tree2 := NewRangeCount(n), NewRangeCount(n)
	tree1.Add(indexMap[nums[0]])
	tree2.Add(indexMap[nums[1]])

	for i := 2; i < len(nums); i++ {
		num := nums[i]
		cnt1 := tree1.CountGT(indexMap[num])
		cnt2 := tree2.CountGT(indexMap[num])
		if cnt1 > cnt2 {
			arr1 = append(arr1, num)
			tree1.Add(indexMap[num])
		} else if cnt1 < cnt2 {
			arr2 = append(arr2, num)
			tree2.Add(indexMap[num])
		} else {
			if len(arr1) <= len(arr2) {
				arr1 = append(arr1, num)
				tree1.Add(indexMap[num])
			} else {
				arr2 = append(arr2, num)
				tree2.Add(indexMap[num])
			}
		}
	}

	// 拼接 arr1 & arr2
	return append(arr1, arr2...)
}

// 暴力模拟法,
// 超时 时间复杂度 O(n^2)
func solution1(nums []int) []int {
	var (
		nums1 []int
		nums2 []int
	)
	nums1 = append(nums1, nums[0])
	nums2 = append(nums2, nums[1])
	for i := 2; i < len(nums); i++ {
		cnt1 := greaterCount(nums1, nums[i])
		cnt2 := greaterCount(nums2, nums[i])

		// if cnt1 == cnt2 && len(nums1) <= len(nums2) || cnt1 > cnt2 {
		// 	nums1 = append(nums1, counts[i])
		// } else {
		// 	nums2 = append(nums2, counts[i])
		// }
		if cnt1 > cnt2 {
			nums1 = append(nums1, nums[i])
		} else if cnt1 < cnt2 {
			nums2 = append(nums2, nums[i])
		} else {
			if len(nums1) <= len(nums2) {
				nums1 = append(nums1, nums[i])
			} else {
				nums2 = append(nums2, nums[i])
			}
		}
	}
	return append(nums1, nums2...)
}

// 返回arr中 严格大于 num 的元素个数
func greaterCount(nums []int, num int) int {
	var cnt int
	for _, item := range nums {
		if item > num {
			cnt++
		}
	}
	return cnt
}
