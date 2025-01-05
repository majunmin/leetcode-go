package leetcode_0729

type interval struct {
	start, end int
}

type intervals []interval

func (iv intervals) Len() int {
	return len(iv)
}

func (iv intervals) Less(i, j int) bool {
	return iv[i].start < iv[j].start
}

func (iv intervals) Swap(i, j int) {
	iv[i], iv[j] = iv[j], iv[i]
}

type MyCalendar struct {
	ivs intervals
}

func Constructor() MyCalendar {
	return MyCalendar{
		ivs: make(intervals, 0),
	}
}

func (this *MyCalendar) Book(startTime int, endTime int) bool {
	if len(this.ivs) == 0 {
		this.ivs = append(this.ivs, interval{
			start: startTime,
			end:   endTime,
		})
		return true
	}
	idx := this.upperBound(endTime)
	if idx == 0 {
		this.ivs = append([]interval{interval{
			start: startTime,
			end:   endTime,
		}}, this.ivs...)
		return true
	}
	if this.ivs[idx-1].end > startTime {
		return false
	}
	// insert
	if idx == len(this.ivs) {
		this.ivs = append(this.ivs, interval{
			start: startTime,
			end:   endTime,
		})
		return true
	}
	this.ivs = append(this.ivs[:idx+1], this.ivs[idx:]...)
	this.ivs[idx] = interval{
		start: startTime,
		end:   endTime,
	}
	return true
}

func (this *MyCalendar) upperBound(val int) int {
	l, r := -1, len(this.ivs)
	for l+1 < r {
		mid := l + (r-l)>>1
		if this.ivs[mid].start < val {
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
 * param_1 := obj.Book(startTime,endTime);
 */
