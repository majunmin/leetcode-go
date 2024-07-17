package lcr_071

import "math/rand/v2"

// https://leetcode.cn/problems/cuyjEf/
type Solution struct {
	preSum []int
	total  int
}

func Constructor(w []int) Solution {
	var (
		preSum = make([]int, 0, len(w))
		total  int
	)
	preSum = append(preSum, w[0])
	total += w[0]
	for i := 1; i < len(w); i++ {
		preSum = append(preSum, w[i]+preSum[i-1])
		total += w[i]
	}
	return Solution{
		preSum: preSum,
		total:  total,
	}
}

func (this *Solution) PickIndex() int {
	//  [1, total]
	randNum := rand.IntN(this.total) + 1
	return this.binarySearch(randNum)
}

// upperBound find >= num的第一个元素
func (this *Solution) binarySearch(num int) int {
	l, r := -1, len(this.preSum)
	for l+1 < r {
		mid := l + (r-l)>>1
		if this.preSum[mid] < num {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
