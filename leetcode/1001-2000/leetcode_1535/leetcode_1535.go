package leetcode_1535

// https://leetcode.cn/problems/find-the-winner-of-an-array-game/?envType=daily-question&envId=2024-05-19
func getWinner(arr []int, k int) int {
	var (
		winCnt = 0
		cur    = arr[0]
	)
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		if num > cur {
			winCnt = 1
			cur = num
		} else {
			winCnt++
			if winCnt == k {
				return cur
			}
		}
	}
	return cur
}
