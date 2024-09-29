package leetcode_0731

var (
	N = 1_000_000_000
)

type Node struct {
	left, right *Node
	val         int
	add         int // 懒标记
}

// https://leetcode.cn/problems/my-calendar-ii/
type MyCalendarTwo struct {
	root *Node
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		root: new(Node),
	}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	if this.query(this.root, 0, N, start, end-1) == 1 {
		return false
	}
	this.update(this.root, 0, N, start, end-1, 1)
	return true
}

func (this *MyCalendarTwo) query(node *Node, left int, right int, start int, end int) int {
	if start <= left && right <= end {
		return node.val
	}
	pushDown(node)
	mid := (left + right) / 2
	result := 0
	if start <= mid {
		result = this.query(node.left, left, mid, start, end)
	}
	if end > mid {
		result = max(result, this.query(node.right, mid+1, right, start, end))
	}
	return result
}

func (this *MyCalendarTwo) update(node *Node, left int, right int, start int, end int, data int) {
	if start <= left && right <= end {
		node.val += data
		node.add += data
		return
	}
	pushDown(node)
	mid := (left + right) / 2
	if end <= mid {
		this.update(node.left, left, mid, start, end, data)
	} else {
		this.update(node.right, mid+1, right, start, end, data)
	}
	pushUp(node)
}

// 每个节点存储当前区间最大值
func pushUp(node *Node) {
	// 每个节点存的是当前区间的最大值
	node.val = max(node.left.val, node.right.val)
}

func pushDown(node *Node) {
	if node.left == nil {
		node.left = new(Node)
	}
	if node.right == nil {
		node.right = new(Node)
	}
	if node.add == 0 {
		return
	}
	node.left.val += node.add
	node.right.val += node.add
	node.left.add += node.add
	node.right.add += node.add
	// 取消当前节点标记
	node.add = 0
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
