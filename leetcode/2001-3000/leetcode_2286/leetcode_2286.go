package leetcode_2286

// https://leetcode.cn/problems/booking-concert-tickets-in-groups/description/
type BookMyShow struct {
	seats  []int
	c      []int // 树状数组
	m      int
	minRow int
}

func Constructor(n int, m int) BookMyShow {
	seats := make([]int, n)
	c := make([]int, n+1)
	for i := 0; i < n; i++ {
		seats[i] = m
		update(c, i+1, m)
	}
	return BookMyShow{
		seats: seats,
		m:     m,
		c:     c,
	}
}

// k > 0
func (this *BookMyShow) Gather(k int, maxRow int) []int {
	for i := this.minRow; i <= maxRow; i++ {
		if this.seats[i] >= k {
			col := this.m - this.seats[i]
			this.seats[i] -= k
			update(this.c, i+1, -k)
			return []int{i, col}
		}
	}
	return []int{}
}

// 可以不连续的进行分配
func (this *BookMyShow) Scatter(k int, maxRow int) bool {
	if q(this.c, maxRow+1) < k {
		return false
	}
	for k > 0 && this.minRow < len(this.seats) {
		seatCnt := this.seats[this.minRow]
		if seatCnt == 0 {
			this.minRow++
			continue
		}
		var leftCnt int
		var decr int
		if seatCnt <= k {
			leftCnt = 0
			k -= seatCnt
			decr = seatCnt
			this.seats[this.minRow] = leftCnt
			update(this.c, this.minRow+1, -decr)
			this.minRow++
		} else {
			leftCnt = seatCnt - k
			decr = k
			k = 0
			this.seats[this.minRow] = leftCnt
			update(this.c, this.minRow+1, -decr)
		}

	}
	return true
}

// 获取最低位的1
func lowbit(x int) int {
	return x & -x
}

func update(arr []int, n, val int) {
	for i := n; i < len(arr); i += lowbit(i) {
		arr[i] += val
	}
}

// 左闭右闭
func q(arr []int, r int) int {
	var result int
	for i := r; i > 0; i -= lowbit(i) {
		result += arr[i]
	}
	return result
}

/**
 * Your BookMyShow object will be instantiated and called as such:
 * obj := Constructor(n, m);
 * param_1 := obj.Gather(k,maxRow);
 * param_2 := obj.Scatter(k,maxRow);
 */
