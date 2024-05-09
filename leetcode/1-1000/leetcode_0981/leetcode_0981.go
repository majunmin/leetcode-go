package leetcode_0981

type entry struct {
	timestamp int
	value     string
}

// https://leetcode.cn/problems/time-based-key-value-store/
type TimeMap struct {
	// key -> ts -> value
	data map[string][]entry
}

func Constructor() TimeMap {
	data := make(map[string][]entry)
	return TimeMap{data: data}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.data[key] = append(this.data[key], entry{
		timestamp: timestamp,
		value:     value,
	})
}

// 1 二分搜索
func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exist := this.data[key]; !exist {
		return ""
	}
	// find 第一个 最后一个 <= timestamp 的值
	idx := lowerBound(this.data[key], 0, len(this.data[key])-1, timestamp+1) - 1
	if idx == -1 {
		return ""
	}
	return this.data[key][idx].value
}

// 红蓝染色法
func lowerBound(entries []entry, left int, right int, target int) int {
	// 开区间搜索
	left, right = left-1, right+1
	for left+1 < right {
		mid := left + (right-left)>>1
		if entries[mid].timestamp >= target {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
