package leetcode_0911

// https://leetcode.cn/problems/online-election/
type TopVotedCandidate struct {
	// list[][0] : 时间
	// list[][1] : 胜利的 人
	list [][]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var (
		winCnt int
		n      = len(persons)
		list   = make([][]int, 0, n)
		votes  = make(map[int]int, n)
	)
	for i := 0; i < len(times); i++ {
		votes[persons[i]] += 1
		if votes[persons[i]] > winCnt {
			winCnt = votes[persons[i]]
			list = append(list, []int{times[i], persons[i]})
		}
	}

	return TopVotedCandidate{list: list}
}

func (this *TopVotedCandidate) Q(t int) int {
	var (
		left, right = -1, len(this.list)
	)
	for left+1 < right {
		mid := left + (right-left)>>1
		if this.list[mid][0] < t+1 {
			left = mid
		} else {
			right = mid
		}
	}
	return this.list[left][1]
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
