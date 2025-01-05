package solution2

import "github.com/emirpasic/gods/trees/redblacktree"

// 差分数组
// 也是解决区间问题的一种思路: https://www.bilibili.com/video/BV1Z2421c7D5/?spm_id_from=333.337.search-card.all.click
//type MyCalendarTwo struct {
//	arr []int
//}
//
//func Constructor() MyCalendarTwo {
//	return MyCalendarTwo{arr: make([]int, 1_000_000_001)}
//}
//
//func (this MyCalendarTwo) add(start, end, val int) {
//	this.arr[start] += val
//	this.arr[end] -= val
//}
//
//func (this MyCalendarTwo) remove(start, end, val int) {
//	this.arr[start] -= val
//	this.arr[end] += val
//}
//
//func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
//	this.add(startTime, endTime, 1)
//	var maxBooks int
//	for i := 0; i < len(this.arr); i++ {
//		maxBooks += this.arr[i]
//		if maxBooks >= 2 {
//			this.remove(startTime, endTime, 1)
//			return false
//		}
//	}
//	return true
//}

type MyCalendarTwo struct {
	rb *redblacktree.Tree
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		rb: redblacktree.NewWithIntComparator(),
	}
}

func (this MyCalendarTwo) add(point, val int) {
	if value, exist := this.rb.Get(point); exist {
		this.rb.Put(point, value.(int)+val)
	} else {
		this.rb.Put(point, val)
	}
}

func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
	this.add(startTime, 1)
	this.add(endTime, -1)
	it := this.rb.Iterator()
	var maxBook int
	for it.Next() {
		maxBook += it.Value().(int)
		if maxBook > 2 {
			this.add(startTime, -1)
			this.add(endTime, 1)
			return false
		}
	}
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
