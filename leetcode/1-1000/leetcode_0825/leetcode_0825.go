package leetcode_0825

import "sort"

// https://leetcode.cn/problems/friends-of-appropriate-ages/?envType=daily-question&envId=2024-11-17
func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	var (
		result int
		prev   int
	)
	ages = append(ages, 130) //方便计算
	for i := 1; i < len(ages); i++ {
		if i > 0 && ages[i] == ages[i-1] {
			continue
		}
		// 统计结果
		//统计 年龄范围(ages[i-1]/2+7,ages[i-1]]的个数
		// result += (right - left +1) * (right-prev)
		left := lower(ages[:i], ages[i-1]/2+7)
		right := i - 1
		if left < right {
			result += (right - left) * (right - prev + 1)
		}
		prev = i
	}
	return result
}

func lower(ages []int, target int) int {
	l, r := -1, len(ages)
	for l+1 < r {
		mid := l + (r-l)>>1
		if ages[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}

//6,6,7,7,7,8,8,8,8,130
//0   6

// 6,7,8,9
