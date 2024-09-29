package lcp_0040

import "sort"

// https://leetcode.cn/problems/uOAnQW/description/?envType=daily-question&envId=2024-08-01
func maxmiumScore(cards []int, cnt int) int {
	if cnt > len(cards) {
		return -1
	}
	var (
		n       = len(cards)
		sum     int
		result  = -1
		minEven = 1001
		minOdd  = 1001
	)
	// 排序后
	sort.Ints(cards)
	for i := n - cnt; i < n; i++ {
		sum += cards[i]
		if cards[i]&1 == 1 {
			minOdd = min(minOdd, cards[i])
		} else {
			minEven = min(minEven, cards[i])
		}
	}
	// 如果为偶数 直接返回.
	if sum&1 == 0 {
		return sum
	}
	//将 cards 从大到小排序后，先贪心的将后 cnt 个数字加起来，若此时 sum 为偶数，直接返回即可。
	//若此时答案为奇数，有两种方案:
	//  在数组前面找到一个最大的奇数与后 cnt 个数中最小的偶数进行替换；
	//  在数组前面找到一个最大的偶数与后 cnt 个数中最小的奇数进行替换。
	//两种方案选最大值即可.
	if minEven != 1001 {
		for i := n - cnt - 1; i >= 0; i-- {
			if cards[i]&1 == 1 {
				result = max(result, sum-minEven+cards[i])
				break
			}
		}
	}
	if minOdd != 1001 {
		for i := n - cnt - 1; i >= 0; i-- {
			if cards[i]&1 == 0 {
				result = max(result, sum-minOdd+cards[i])
				break
			}
		}
	}
	return result
}
