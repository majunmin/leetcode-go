package solution1

type interval struct {
	start, end int
}

type MyCalendarTwo struct {
	booked   []interval
	overlaps []interval
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		booked:   make([]interval, 0),
		overlaps: make([]interval, 0),
	}
}

func (this *MyCalendarTwo) Book(startTime int, endTime int) bool {
	for _, item := range this.overlaps {
		if item.start < endTime || item.end > startTime {
			return false
		}
	}
	// 可以进行预订.
	for _, item := range this.booked {
		if startTime < item.end && item.start < endTime {
			this.overlaps = append(this.overlaps, interval{
				start: max(startTime, item.start),
				end:   min(endTime, item.end),
			})
		}
	}
	this.booked = append(this.booked, interval{
		start: startTime,
		end:   endTime,
	})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
