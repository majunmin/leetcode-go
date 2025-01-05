package solution1

// 内存溢出.

type Node struct {
	l, r int //闭区间
	mid  int
	val  int

	left  *Node // l child
	right *Node // r child
}

func newNode(l, r int) *Node {
	return &Node{
		l:   l,
		mid: l + (r-l)/2,
		r:   r,
		val: 0,
	}
}

func (n *Node) add(l int, r int, val int) {
	// terminate
	if l > r {
		return
	}
	// 不存在交集
	if n.l > r || n.r < l {
		return
	}

	if n.l == n.r {
		n.val += val
		return
	}
	if l <= n.mid {
		if n.left == nil {
			n.left = newNode(n.l, n.mid)
		}
		n.left.add(l, min(r, n.mid), val)
	}
	if r > n.mid {
		if n.right == nil {
			n.right = newNode(n.mid+1, n.r)
		}
		n.right.add(max(n.mid+1, l), r, val)
	}
}

func (n *Node) query() int {
	val := n.val
	if n.left != nil {
		val = max(val, n.left.query())
	}
	if n.right != nil {
		val = max(val, n.right.query())
	}
	return val
}

type segmentTree struct {
	root *Node
}

func newSegmentTree(l, r int) *segmentTree {
	return &segmentTree{
		root: newNode(l, r),
	}
}

func (st *segmentTree) add(l, r, val int) {
	st.root.add(l, r-1, val)
}

func (st *segmentTree) query(l, r int) int {
	return st.root.query()
}

// https://leetcode.cn/problems/my-calendar-iii/?envType=daily-question&envId=2025-01-04
type MyCalendarThree struct {
	segmentTree *segmentTree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		segmentTree: newSegmentTree(0, 1_000_000_000),
	}
}

func (this *MyCalendarThree) Book(startTime int, endTime int) int {
	this.segmentTree.add(startTime, endTime, 1)
	return this.segmentTree.query(0, 1_000_000_000)
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
