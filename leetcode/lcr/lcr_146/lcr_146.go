package lcr_146

func spiralArray(array [][]int) []int {
	var (
		m, n   = len(array), len(array[0])
		l, r   = 0, n - 1
		t, b   = 0, m - 1
		result = make([]int, m*n)
	)
	for l <= r || t <= b {
		// 从左向右
		for i := l; i <= r; i++ {
			result = append(result, array[t][i])
		}
		t++
		if t > b {
			break
		}
		// 从上至下
		for i := t; i <= b; i++ {
			result = append(result, array[i][r])
		}
		r--
		if l > r {
			break
		}
		// 从右至左
		for i := r; i >= l; i-- {
			result = append(result, array[b][i])
		}
		b--
		if t > b {
			break
		}
		// 从下至上
		for i := b; i >= t; i-- {
			result = append(result, array[i][l])
		}
		l++
	}
	return result
}
