package solution2

// https://leetcode.cn/problems/snapshot-array/?envType=daily-question&envId=2024-04-26
// 超出内存限制
type SnapshotArray struct {
	length  int
	snapIdx int
	history [][][2]int
	current []int // 题目要求来说 不是必须的.
}

func Constructor(length int) SnapshotArray {
	history := make([][][2]int, length)
	for i := range history {
		history[i] = append(history[i], [2]int{0, 0})
	}
	return SnapshotArray{
		length:  length,
		history: history,
		current: make([]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.current[index] = val
	this.history[index] = append(this.history[index], [2]int{this.snapIdx, val})
}

func (this *SnapshotArray) Snap() int {
	this.snapIdx++
	return this.snapIdx - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// param check
	if index < 0 || index >= this.length {
		return 0
	}
	versions := this.history[index]
	// 二分查找 找到 <= snap_id 的snap 的最后一次修改.
	l, r := 0, len(versions)-1
	for l < r {
		mid := l + (r-l+1)>>1
		if snap_id < versions[mid][0] {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return versions[l][1]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
