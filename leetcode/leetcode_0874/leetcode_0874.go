package leetcode_0874

type node struct {
	x int
	y int
}

// https://leetcode.cn/problems/walking-robot-simulation/description/
func robotSim(commands []int, obstacles [][]int) int {
	//if len(commands) == 0  {
	//	return 0
	//}
	directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	var mp map[node]struct{}
	for _, obs := range obstacles {
		mp[node{obs[0], obs[1]}] = struct{}{}
	}

	x, y, di := 0, 0, 0
	result := 0

	for _, cmd := range commands {
		// 循环数组的移动
		if cmd == -2 {
			di = (di + 3) % 4
			continue
		}
		if cmd == -1 {
			di = (di + 1) % 4
			continue
		}
		// 走路
		for i := 0; i < cmd; i++ {
			tx, ty := x+directions[di][0], y+directions[di][1]
			// 遇到 障碍物
			if _, exist := mp[node{tx, ty}]; exist {
				break
			}
			x, y = tx, ty

		}
		if dis := x*x + y*y; dis > result {
			result = dis
		}

	}
	return result
}
