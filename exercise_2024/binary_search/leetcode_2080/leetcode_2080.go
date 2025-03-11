package leetcode_2080

// https://leetcode.cn/problems/range-frequency-queries/
type RangeFreqQuery struct {
	n              int
	item2IdxSlices map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	item2IdxSlices := make(map[int][]int, len(arr))
	for i, item := range arr {
		item2IdxSlices[item] = append(item2IdxSlices[item], i)
	}
	return RangeFreqQuery{
		n:              len(arr),
		item2IdxSlices: item2IdxSlices,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	//param check
	if left < 0 || right > this.n || left > right {
		return -1
	}
	//
	l := lowerBound(this.item2IdxSlices[value], left)
	r := lowerBound(this.item2IdxSlices[value], right+1) - 1
	return r - l + 1
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

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
