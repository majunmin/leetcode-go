package leetcode_0729

type interval struct {
	start, end int
}

// https://leetcode.cn/problems/my-calendar-i/
type MyCalendar struct {
	intervals []interval
}

func Constructor() MyCalendar {
	return MyCalendar{
		intervals: make([]interval, 0, 1000),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	if len(this.intervals) == 0 {
		this.insert(0, start, end)
		return true
	}
	idx := this.upperBound(end)
	if idx == 0 {
		this.insert(idx, start, end)
		return true
	} else if this.intervals[idx-1].end <= start {
		this.insert(idx, start, end)
		return true
	}
	return false
}

func (this *MyCalendar) insert(i, start int, end int) {
	if i == len(this.intervals) {
		this.intervals = append(this.intervals, interval{
			start: start,
			end:   end,
		})
	} else {
		this.intervals = append(this.intervals[:i+1], this.intervals[i:]...)
		this.intervals[i].start = start
		this.intervals[i].end = end
	}
}

func (this *MyCalendar) upperBound(end int) int {
	l, r := -1, len(this.intervals)
	for l+1 < r {
		mid := l + (r-l)>>1
		if this.intervals[mid].start < end {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
