package leetcode_0134

// https://leetcode.cn/problems/gas-station/?envType=study-plan-v2&envId=top-interview-150
func canCompleteCircuit(gas []int, cost []int) int {
	// todo: ????
	var (
		n = len(gas)
	)

	// 模拟
	for i := 0; i < n; i++ {
		if gas[i] < cost[i] {
			continue
		}
		left := gas[i]
		var notDone bool
		for j := 0; j < n; j++ {
			cur := (i + j + n) % n
			next := (i + j + n + 1) % n
			left -= cost[cur]
			if left < 0 {
				notDone = true
				break
			}
			left += gas[next]
		}
		if !notDone {
			return i
		}
	}
	return -1
}
