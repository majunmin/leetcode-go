package leetcode_0860

// https://leetcode.cn/problems/lemonade-change/description/
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}
	// 钱包 包含 5, 10, 20 钞票个数
	bag := make([]int, 3)
	for _, bill := range bills {
		if bill-5 == 0 {
			bag[0]++
			continue
		}
		if bill-5 == 5 {
			if bills[0] == 0 {
				return false
			}
			bag[0]--
			bag[1]++
			continue
		}
		if bill-5 == 15 {
			if bag[1] > 0 && bag[0] > 0 {
				bag[1]--
				bag[0]--
				bag[2]++
			} else if bag[0] >= 3 {
				bag[0] = bag[0] - 3
				bag[2]++
			} else {
				return false
			}
			continue
		}
		return false
	}
	return true
}
