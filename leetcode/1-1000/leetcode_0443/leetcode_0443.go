package leetcode_0443

import "math"

var (
	minDepth = math.MaxInt
)

// https://leetcode.cn/problems/minimum-genetic-mutation/#/description
func minMutation(start string, end string, bank []string) int {
	visited := make(map[string]bool, len(bank))
	dfs(start, end, bank, visited, 0)
	if minDepth == math.MaxInt {
		return -1
	}
	return minDepth
}

func dfs(start string, end string, bank []string, visited map[string]bool, step int) {
	if start == end {
		if step < minDepth {
			minDepth = step
		}
	}

	// for choice in choiceList
	for _, choice := range bank {
		if !visited[choice] && diff(choice, start) == 1 {
			visited[choice] = true
			dfs(choice, end, bank, visited, step+1)
			visited[choice] = false
		}
	}
}

func diff(str1 string, str2 string) int {
	cnt := 0
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			cnt++
		}
	}
	return cnt
}

// 双向 BFS解法
func doubleBfsSolution(start string, end string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}
	// start end 不能相等
	if start == end {
		return 0
	}

	set := make(map[string]struct{})
	for _, str := range bank {
		if str == start {
			continue
		}
		set[str] = struct{}{}
	}
	if _, exist := set[end]; !exist {
		return -1
	}

	table := []rune{'A', 'T', 'C', 'G'}
	startQueue := []string{start}
	endQueue := []string{end}
	step := 0
	for len(startQueue) > 0 && len(endQueue) > 0 {
		var newQueue []string
		for _, curStr := range startQueue {
			for i, m := range curStr {
				for _, n := range table {
					if m == n {
						continue
					}
					next := curStr[:i] + string(n) + curStr[i+1:]
					if _, exist := set[next]; !exist {
						continue
					}
					for _, endStr := range endQueue {
						if next == endStr {
							return step + 1
						}
					}
					newQueue = append(newQueue, next)
				}
			}

		}
		step++
		startQueue = newQueue
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return -1
}

// BFS 解法
func bfsSolution(start string, end string, bank []string) int {
	if len(bank) == 0 {
		return -1
	}
	// start end 不能相等
	if start == end {
		return 0
	}

	set := make(map[string]struct{})
	for _, str := range bank {
		set[str] = struct{}{}
	}
	// 如果 end 不在基因库中, 返回 false
	if _, exist := set[end]; !exist {
		return -1
	}

	step := 0
	var queue []string
	queue = append(queue, start)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			curStr := queue[i]
			for j, x := range curStr {
				for _, y := range "ACGT" {
					if x == y {
						continue
					}
					next := curStr[:j] + string(y) + curStr[j+1:]
					if _, exist := set[next]; !exist {
						continue
					}
					if next == end {
						return step + 1
					}
					queue = append(queue, next)
				}
			}
			delete(set, curStr)
		}
		queue = queue[l:]
		step++
	}
	return -1
}
