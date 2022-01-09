// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-09 12:38
package leetcode

//https://leetcode-cn.com/problems/slowest-key/
func slowestKey(releaseTimes []int, keysPressed string) byte {
	return iterSolution(releaseTimes, keysPressed)
}

// 迭代  keysPressed
func iterSolution(releaseTimes []int, keysPressed string) byte {
	if len(releaseTimes) == 0 || len(releaseTimes) != len(keysPressed) {
		panic("param error.")
	}
	res := keysPressed[0]
	maxTime := releaseTimes[0]

	for i := 1; i < len(keysPressed); i++ {
		k := keysPressed[i]
		time := releaseTimes[i] - releaseTimes[i-1]
		if time > maxTime || time == maxTime && k > res {
			res = k
			maxTime = time
		}
	}
	return res
}
