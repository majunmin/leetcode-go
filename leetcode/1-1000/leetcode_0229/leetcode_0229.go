package leetcode_0229

// https://leetcode.cn/problems/majority-element-ii/
func majorityElement(nums []int) []int {
	// param check
	if len(nums) == 0 {
		return nil
	}
	var (
		candiate1 int
		vote1     int
		candiate2 int
		vote2     int
	)
	for _, num := range nums {
		if candiate1 == num {
			vote1++
			continue
		}
		if candiate2 == num {
			vote2++
			continue
		}
		// 如果都没有匹配到
		if vote1 == 0 {
			candiate1 = num
			vote1 = 1
			continue
		}
		if vote2 == 0 {
			candiate2 = num
			vote2 = 1
			continue
		}
		// 如果 vote1 > 0 && vote2 > 0
		vote1--
		vote2--
	}
	// 2. 计数阶段
	var (
		count1 int
		count2 int
		result []int
	)

	for _, num := range nums {
		if num == candiate1 {
			count1++
			// 这里的目的是去掉重复数字
		} else if num == candiate2 {
			count2++
		}
	}
	if count1 > len(nums)/3 {
		result = append(result, candiate1)
	}
	if count2 > len(nums)/3 {
		result = append(result, candiate2)
	}
	return result
}
