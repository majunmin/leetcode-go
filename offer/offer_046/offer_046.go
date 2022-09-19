package offer_046

import "strconv"

//https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/?envType=study-plan&id=lcof
// 类似青蛙跳台
func translateNum(num int) int {
	str := strconv.Itoa(num)
	n := len(str)
	p, q, r := 0, 1, 1
	for i := 1; i < n; i++ {
		p = q
		q = r
		pre := str[i-1 : i+1]
		if "10" <= pre && pre <= "25" {
			r += p
		}
	}
	return r

}

func recurSolution(num int) int {
	if num < 10 {
		return 1
	}

	a := translateNum(num / 10)

	b := 0
	two := num % 100
	if 10 <= two && two < 25 {
		b = translateNum(num / 100)
	}

	// f(n) = f(n/10) + f(n/100)
	return a + b
}
