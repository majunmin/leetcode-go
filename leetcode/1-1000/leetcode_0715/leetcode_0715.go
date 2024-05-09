package leetcode_0715

import "github.com/emirpasic/gods/trees/redblacktree"

type RangeModule struct {
	*redblacktree.Tree
}

func Constructor() RangeModule {
	return RangeModule{
		redblacktree.NewWithIntComparator(),
	}
}

func (rm *RangeModule) AddRange(left int, right int) {
	node, exist := rm.Floor(left)
	if exist {
		// node.Left <= left
		l := node.Key.(int)
		r := node.Value.(int)
		if r >= right {
			return
		}
		if r >= left {
			left = l
			rm.Remove(l)
		}
		// 确定了 左边界👈🏻
	}
	// 确定右边界
	for node, exist := rm.Ceiling(left); exist; {
		l, r := node.Key.(int), node.Value.(int)
		if l > right {
			break
		}
		right = max(right, r)
		rm.Remove(l)
	}
	rm.Put(left, right)
}

func (rm *RangeModule) QueryRange(left int, right int) bool {
	node, exist := rm.Floor(left)
	return exist && right <= node.Value.(int)
}

func (rm *RangeModule) RemoveRange(left int, right int) {
	node, exist := rm.Floor(left)
	if exist {
		l, r := node.Key.(int), node.Value.(int)
		if right <= r {
			if l == left {
				rm.Remove(l)
			} else {
				node.Value = left
			}
			if right != r {
				rm.Put(right, r)
			}
			return
		}
		if r > left {
			if l == left {
				rm.Remove(l)
			} else {
				node.Value = left
			}
		}
	}

	for node, exist = rm.Ceiling(left); exist; {
		l, r := node.Key.(int), node.Value.(int)
		if l >= right {
			break
		}
		rm.Remove(l)
		if r > right {
			rm.Put(right, r)
			break
		}
		// iter
		node, exist = rm.Ceiling(left)
	}

}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
