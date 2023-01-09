package offer_029

// https://leetcode.cn/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/?envType=study-plan&id=lcof
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	t, b := 0, len(matrix)-1    // top, bottom
	l, r := 0, len(matrix[0])-1 // left, right
	result := make([]int, len(matrix)*len(matrix[0]))
	idx := 0
	for {
		// ->
		for i := l; i <= r; i++ {
			result[idx] = matrix[t][i]
			idx++
		}
		t++
		if t > b {
			break
		}
		// |
		// v
		for i := t; i <= b; i++ {
			result[idx] = matrix[i][r]
			idx++
		}
		r--
		if l > r {
			break
		}
		// <-
		for i := r; i >= l; i-- {
			result[idx] = matrix[b][i]
			idx++
		}
		b--
		if t > b {
			break
		}

		// ^
		// |
		for i := b; i >= t; i-- {
			result[idx] = matrix[i][l]
			idx++
		}
		l++
		if l > r {
			break
		}
	}
	return result
}
