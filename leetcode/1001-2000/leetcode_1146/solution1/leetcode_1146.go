package solution1

// https://leetcode.cn/problems/snapshot-array/?envType=daily-question&envId=2024-04-26
// 超出内存限制
type SnapshotArray struct {
	//data      []int
	current   []int
	snapshots [][]int
	length    int
	latestIdx int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		length:    length,
		current:   make([]int, length),
		snapshots: make([][]int, 0),
		latestIdx: -1,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	this.current[index] = val
}

func (this *SnapshotArray) Snap() int {
	snapshot := make([]int, this.length)
	copy(snapshot, this.current)
	this.snapshots = append(this.snapshots, snapshot)
	this.latestIdx++
	return this.latestIdx
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// param check
	return this.snapshots[snap_id][index]
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
