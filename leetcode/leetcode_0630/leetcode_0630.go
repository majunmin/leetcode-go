package leetcode_0630

import (
	"container/heap"
	"sort"
)

type PriorityQueue [][]int

func (h *PriorityQueue) Len() int {
	return len(*h)
}

func (h *PriorityQueue) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *PriorityQueue) Less(i, j int) bool {
	return (*h)[i][0] > (*h)[j][0]
}

func (h *PriorityQueue) Push(x any) {
	item := x.([]int)
	*h = append(*h, item)
}

func (h *PriorityQueue) Pop() any {
	item := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return item
}

// https://leetcode.cn/problems/course-schedule-iii/
func scheduleCourse(courses [][]int) int {
	// [t1, d1], [t2, d2],[t3, d3],[t4, d4]
	// t1 < d1 , 第一门课可以加进去
	// t1 + t2 < d2 , 第二门课可以加进去
	// t1 + t2 + t3 < d3 , 第三门课可以加进去
	//贪心算法

	// 1. param check

	// 2. 按照 deadline 排序
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	var pq PriorityQueue
	heap.Init(&pq)
	// 3. 用大顶堆 存储  课程持续时间
	var (
		sum int
	)
	for _, course := range courses {
		// key加入 该课程
		heap.Push(&pq, course)
		sum += course[0]
		if sum > course[1] {
			x := heap.Pop(&pq)
			item := x.([]int)
			sum -= item[0]
		}
	}

	return len(pq)
}
