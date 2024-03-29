package leetcode_0022

func generateParenthesis(n int) []string {
	// param check
	if n < 0 {
		panic("invalid param")
	}
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	resultMap := make(map[int][]string)
	resultMap[0] = []string{""}
	resultMap[1] = []string{"()"}
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, left := range resultMap[j] {
				for _, right := range resultMap[i-j-1] {
					resultMap[i] = append(resultMap[i], "("+left+")"+right)
				}
			}
		}
	}

	return resultMap[n]
}

func backtraceSolution(n int) []string {
	var result []string
	backtrace(&result, n, 0, 0, "")
	return result
}

func backtrace(result *[]string, n int, l, r int, path string) {
	// terminate
	if l+r == n*2 {
		temp := make([]byte, n)
		copy(temp, path)
		*result = append(*result, string(temp))
		return
	}
	// for choice in choiceList
	if l < n {
		backtrace(result, n, l+1, r, path+"(")
	}
	if r < l {
		backtrace(result, n, l, r+1, path+")")
	}
}
