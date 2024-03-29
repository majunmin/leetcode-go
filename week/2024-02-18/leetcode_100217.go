package _024_02_18

import (
	"container/heap"
	"math"
)

var (
	directions = [][]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}
)

type item struct {
	val int
	// 没有业务含义,为了方便更新操作而增加的值, 堆中每个元素都需要携带
	index     int
	frequency int
}

type maxHeap struct {
	items   []*item
	itemMap map[int]*item
}

func newMaxHeap() *maxHeap {
	mHp := &maxHeap{
		items:   make([]*item, 0),
		itemMap: make(map[int]*item),
	}
	heap.Init(mHp)
	return mHp
}

func (m *maxHeap) Len() int {
	return len(m.items)
}

func (m *maxHeap) Less(i, j int) bool {
	if m.items[i].frequency > m.items[j].frequency {
		return true
	} else if m.items[i].frequency < m.items[j].frequency {
		return false
	} else {
		return m.items[i].val > m.items[j].val
	}
}

func (m *maxHeap) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
	m.items[i].index = i
	m.items[j].index = j
}

func (m *maxHeap) Push(x any) {
	val := x.(int)
	if _, exist := m.itemMap[val]; exist {
		m.addFrequency(m.itemMap[val])
	} else {
		index := len(m.items)
		v := &item{
			val:       val,
			index:     index,
			frequency: 1,
		}
		m.items = append(m.items, v)
		m.itemMap[val] = v
	}
}

func (m *maxHeap) Pop() any {
	v := m.items[len(m.items)-1]
	m.items = m.items[:len(m.items)-1]
	return v
}

func (m *maxHeap) push(x any) {
	heap.Push(m, x)
}

func (m *maxHeap) pop() any {
	return heap.Pop(m)
}

func (m *maxHeap) addFrequency(item *item) {
	item.frequency++
	heap.Fix(m, item.index)
}

// https://leetcode.cn/contest/weekly-contest-385/problems/most-frequent-prime/
func mostFrequentPrime(mat [][]int) int {
	// param check
	if len(mat) == 0 || len(mat[0]) == 0 {
		return -1
	}
	m, n := len(mat), len(mat[0])
	// visited map
	mHeap := newMaxHeap()
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, dir := range directions {
				dfs(mat, visited, mHeap, i, j, 0, dir)
			}

		}
	}
	if mHeap.Len() > 0 {
		node := mHeap.pop().(*item)
		return node.val
	}
	return -1
}

func dfs(mat [][]int, visited [][]bool, mHeap *maxHeap, i int, j int, curVal int, dir []int) {
	curVal = curVal*10 + mat[i][j]
	if curVal > 10 && isSushu(curVal) {
		mHeap.push(curVal)
	}
	visited[i][j] = true
	// for choice in choiceList
	newx, newy := i+dir[0], j+dir[1]
	if newx >= 0 && newy >= 0 && newx < len(mat) && newy < len(mat[0]) && !visited[newx][newy] {
		dfs(mat, visited, mHeap, newx, newy, curVal, dir)
	}
	visited[i][j] = false
}

func isSushu(num int) bool {
	if num == 1 { // 1不是质数
		return false
	}
	for i := 2; i < int(math.Sqrt(float64(num)))+1; i++ { // 直接开根号，能让这个数的范围减少
		if num%i == 0 { // 如果能有数被整出，那么就不是素数了
			return false
		}
	}
	return true
}
