package offer_066

// https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/?envType=study-plan&id=lcof
func constructArr(a []int) []int {
	// 利用前缀和原理
	// 左右乘积前缀
	length := len(a)
	if length == 0 {
		return a[:]
	}

	l, r := make([]int, length), make([]int, length)
	// l[i] :前缀积， index = i 左边的数之积(不包含 i)
	// r[i] :后缀积， index = i 右边的数之积(不包含 i)
	// 那么 result[i] = l[i] * r[i]
	l[0] = 1
	for i := 1; i < length; i++ {
		l[i] = l[i-1] * a[i-1]
	}
	r[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		r[i] = r[i+1] * a[i+1]
	}

	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = l[i] * r[i]
	}
	return result

}

// 注意 数组元素为0 时处理
func errSolution(a []int) []int {
	if len(a) == 0 {
		return a[:]
	}
	target := 1
	for i := 0; i < len(a); i++ {
		target *= a[i]
	}

	result := make([]int, len(a))
	for i := 0; i < len(result); i++ {
		result[i] = target / a[i]
	}
	return result
}
