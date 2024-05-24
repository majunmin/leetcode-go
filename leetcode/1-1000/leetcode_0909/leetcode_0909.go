package leetcode_0909

// https://leetcode.cn/problems/snakes-and-ladders/description/?envType=study-plan-v2&envId=top-interview-150
func snakesAndLadders(board [][]int) int {
	var (
		n     = len(board)
		total = n * n
		queue = make([]int, 0)
		memo  = make(map[int]int)
		step  int
	)
	queue = append(queue, 0)
	memo[0] = 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			if item == total {
				return memo[total]
			}
			for j := 1; j <= 6; j++ {
				next := item + j
				if next > total {
					break
				}
				rc := id2rc(next, n)
				if board[rc[0]][rc[1]] != -1 {
					next = board[rc[0]][rc[1]]
				}
				if memo[next] > 0 {
					continue
				}
				memo[next] = step + 1
				queue = append(queue, next)
			}
		}
		step++
		queue = queue[size:]
	}
	return -1
}

func id2rc(id, n int) []int {
	r := (id - 1) / n
	c := (id - 1) % n
	if r&1 == 1 {
		c = n - 1 - c
	}
	r = n - 1 - r
	return []int{r, c}
}
