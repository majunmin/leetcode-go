// Created By: junmin.ma
// Description: <description>
// Date: 2022-01-10 21:18
package leetcode_0306

// 字符串加法
func stringAdd(s1 string, s2 string) string {
	var res []byte

	carry, cur := 0, 0
	for s1 != "" || s2 != "" || carry != 0 {
		cur = carry
		if s1 != "" {
			cur += int(s1[len(s1)-1] - '0')
			s1 = s1[:len(s1)-1]
		}
		if s2 != "" {
			cur += int(s2[len(s2)-1] - '0')
			s2 = s2[:len(s2)-1]
		}

		carry = cur / 10
		res = append(res, byte(cur%10)+'0')
	}

	// 倒序
	for i, n := 0, len(res)-1; i < len(res)/2; i++ {
		res[i], res[n-i] = res[n-i], res[i]
	}
	return string(res)
}

func isValid(num string, secondStart int, secondEnd int) bool {
	n := len(num)
	firstStart, firstEnd := 0, secondStart-1
	for secondStart <= n-1 {
		third := stringAdd(num[firstStart:firstEnd+1], num[secondStart:secondEnd+1])
		thirdStart := secondEnd + 1
		thirdEnd := thirdStart + len(third) - 1

		// terminate
		if thirdEnd >= n || num[thirdStart:thirdEnd+1] != third {
			break
		}

		if thirdEnd == n-1 {
			return true
		}

		firstStart, firstEnd = secondStart, secondEnd
		secondStart, secondEnd = thirdStart, thirdEnd
	}
	return false

}

// 判断是否是累加数
// https://leetcode-cn.com/problems/additive-number/
func isAdditiveNumber(num string) bool {
	n := len(num)

	for secondStart := 1; secondStart < n-1; secondStart++ {
		if num[0] == '0' && secondStart != 1 {
			break
		}

		for secondEnd := secondStart; secondEnd < n-1; secondEnd++ {
			if num[secondStart] == '0' && secondStart != secondEnd {
				break
			}
			if isValid(num, secondStart, secondEnd) {
				return true
			}
		}
	}
	return false
}
