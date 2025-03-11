package leetcode_1170

import "slices"

//

// https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/
func numSmallerByFrequency(queries []string, words []string) []int {
	return solution2(queries, words)
}

// 根据数据量进行优化
// https://leetcode.cn/problems/compare-strings-by-frequency-of-the-smallest-character/solutions/2297291/bi-jiao-zi-fu-chuan-zui-xiao-zi-mu-chu-x-pb50/
func solution2(queries []string, words []string) []int {
	//  len(word[i])最长为 10,
	count := make([]int, 12)
	for _, item := range words {
		count[calMinCharFreq(item)]++
	}
	for i := 9; i >= 0; i-- {
		count[i] += count[i+1]
	}

	result := make([]int, 0, len(queries))
	for _, item := range queries {
		num := calMinCharFreq(item)
		result = append(result, count[num+1])
	}
	return result
}

// 版本1:
// 1. 计算单词最小字母频次 calMinCharFreq
// 2. 将 频次 记录到一个数组并从小到大排序.
// 3. 二分查找 >= f(query[i])的索引 idx,那么符合条件的单词数是 len(words) - idx.
func solution1(queries []string, words []string) []int {
	// param check
	if len(queries) == 0 {
		return nil
	}
	var (
		result = make([]int, 0, len(queries))
		nums   = make([]int, 0, len(words))
	)
	for _, item := range words {
		nums = append(nums, calMinCharFreq(item))
	}
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})
	for _, item := range queries {
		freq := calMinCharFreq(item)
		val := len(words) - lowerBound(nums, freq+1)
		result = append(result, val)
	}
	return result
}

func lowerBound(nums []int, target int) int {
	l, r := -1, len(nums)
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

func calMinCharFreq(item string) int {
	var (
		freq int
		minC = 'z'
	)
	for _, c := range item {
		if c > minC {
			continue
		}
		if c < minC {
			freq = 1
			minC = c
		} else {
			// c == minC
			freq++
		}

	}
	return freq
}
