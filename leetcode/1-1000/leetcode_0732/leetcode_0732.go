package leetcode_0732

import "github.com/emirpasic/gods/trees/redblacktree"

type MyCalendarThree struct {
	rb *redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		rb: redblacktree.NewWithIntComparator(),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.add(startTime, 1)
	this.add(endTime-1, -1)
	iter := this.rb.Iterator()
	var maxCnt int
	var cur int
	if iter.Next() {
		cur += iter.Value().(int)
		maxCnt = max(maxCnt, cur)
	}
	return maxCnt
}

func (this *MyCalendarThree) add(startTime int, val int) {
	var delta int
	value, exist := this.rb.Get(startTime)
	if exist {
		delta = value.(int)
	}
	this.rb.Put(startTime, delta+val)
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
