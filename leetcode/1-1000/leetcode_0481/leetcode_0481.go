package leetcode_0481

// https://leetcode.cn/problems/magical-string/
func magicalString(n int) int {
	// param check
	if n <= 0 {
		return -1
	}
	// 1 2 2...
	if n < 4 {
		return 1
	}

	// init state
	s := make([]int, n)
	s[0], s[1], s[2] = 1, 2, 2
	cnt := 1
	i, j := 2, 3
	for j < n {
		// repeat count
		item := 3 - s[j-1]
		size := s[i]
		for size > 0 && j < n {
			s[j] = item
			j++
			// 计数
			if item == 1 {
				cnt++
			}
			size--
		}
		i++
	}

	return cnt
}
