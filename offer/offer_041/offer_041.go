package offer_041

// https://leetcode.cn/problems/shu-ju-liu-zhong-de-zhong-wei-shu-lcof/?envType=study-plan&id=lcof
type MedianFinder struct {
	arr    []int
	length int
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		arr:    make([]int, 0, 16),
		length: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {

	for i := 0; i < this.length; i++ {
		if this.arr[i] > num {
			this.arr = append(this.arr[:i], append([]int{num}, this.arr[i:]...)...)
			this.length++
			return
		}
	}

	this.arr = append(this.arr, num)
	this.length++
}

func (this *MedianFinder) FindMedian() float64 {
	if this.length == 0 {
		return 0
	}
	mid := this.length / 2
	if this.length&1 == 1 { //奇数
		return float64(this.arr[mid])
	} else {
		return float64(this.arr[mid]+this.arr[mid-1]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
