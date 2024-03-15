package leetcode_2575

// https://leetcode.cn/problems/find-the-divisibility-array-of-a-string/description/
func divisibilityArray(word string, m int) []int {
	//
	var (
		curNum int
		result = make([]int, len(word))
	)

	for i := range word {
		curNum = (curNum*10 + int(word[i]-'0')) % m
		if curNum%m == 0 {
			result[i] = 1
		}
	}
	return result
}

// 一个整数可以表示为 a * 10 + b
// 那么 (a * 10 + b) mod m = ((a mod m) * 10 + b) mod m

// 假设 a = k1m +r1 , b = k2m + r2
// (a + b)mod m == (r1 + r2) mod m == (a mod m + b mod m) mod m
// 这意味着 我们可以在计算中取模计算,而不是 到了最后在计算
