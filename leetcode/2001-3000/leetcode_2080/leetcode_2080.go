package leetcode_2080

// https://leetcode.cn/problems/range-frequency-queries/
type RangeFreqQuery struct {
	num2Idx [][]int
}

func Constructor(arr []int) RangeFreqQuery {
	num2Idx := make([][]int, 10_000)
	// 索引是有序的
	for i, num := range arr {
		num2Idx[num] = append(num2Idx[num], i)
	}
	return RangeFreqQuery{num2Idx: num2Idx}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	if len(this.num2Idx[value]) == 0 {
		return 0
	}
	leftIdx := lowerBound(this.num2Idx[value], left)
	rightIdx := lowerBound(this.num2Idx[value], (right+1)) - 1
	return rightIdx - leftIdx + 1
}

func lowerBound(idxs []int, target int) int {
	var (
		left, right = -1, len(idxs)
	)
	for left+1 < right {
		mid := left + (right-left)>>1
		if idxs[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
