package leetcode_0013

var (
	roman2Int = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

// https://leetcode.cn/problems/roman-to-integer/?envType=study-plan-v2&envId=top-interview-150
func romanToInt(s string) int {
	n := len(s)
	var result int
	for i := 0; i < n; i++ {
		val := roman2Int[s[i]]
		if i < n-1 && val < roman2Int[s[i+1]] {
			result -= val
		} else {
			result += val
		}
	}
	return result
}
