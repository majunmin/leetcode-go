package leetcode_0715_2

import (
	"cmp"
	"slices"
)

type intervals []*interval

type interval struct {
	start int
	end   int
}

type RangeModule struct {
	data intervals
}

func Constructor() RangeModule {
	return RangeModule{
		make(intervals, 0),
	}
}

func (rm *RangeModule) AddRange(left int, right int) {
	newIntervals := make(intervals, 0, len(rm.data))
	var i int
	for i < len(rm.data) && left > rm.data[i].end {
		newIntervals = append(newIntervals, rm.data[i])
		i++
	}
	// left <= rm.data[i].end
	for i < len(rm.data) && rm.data[i].start <= right {
		left = min(left, rm.data[i].start)
		right = max(right, rm.data[i].end)
	}
	newIntervals = append(newIntervals, &interval{left, right})
	newIntervals = append(newIntervals, rm.data[i:]...)
	rm.data = newIntervals
}

func (rm *RangeModule) QueryRange(left int, right int) bool {
	// 可以使用 二分查找算法 进行快速查询
	res, exist := slices.BinarySearchFunc(rm.data, &interval{left, right}, func(i *interval, j *interval) int {
		return cmp.Compare(i.start, j.start)
	})
	return exist && right <= rm.data[res].end
}

func (rm *RangeModule) RemoveRange(left int, right int) {
	newIntervals := make(intervals, 0, len(rm.data))
	var i int
	for i < len(rm.data) && left >= rm.data[i].end {
		newIntervals = append(newIntervals, rm.data[i])
	}
	// left < rm.data[i].end
	for i < len(rm.data) && rm.data[i].start < right {
		if rm.data[i].start < left {
			newIntervals = append(newIntervals, &interval{rm.data[i].start, left})
		}
		if rm.data[i].end > right {
			newIntervals = append(newIntervals, &interval{right, rm.data[i].end})
		}
		i++
	}
	newIntervals = append(newIntervals, rm.data[i:]...)
	rm.data = newIntervals
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
