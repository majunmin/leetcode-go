package leetcode_0826

// https://leetcode.cn/problems/most-profit-assigning-work/
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	sortDifficulty(difficulty, profit)
	var (
		idx          int
		curMaxPorfit int
		result       int
	)

	for _, cap := range worker {
		for idx < len(difficulty) && difficulty[idx] <= cap {
			curMaxPorfit = max(curMaxPorfit, profit[idx])
			idx++
		}
		result += curMaxPorfit
	}
	return result
}

func sortDifficulty(difficulty []int, profit []int) {
	quickSort(difficulty, profit, 0, len(difficulty)-1)
}

func quickSort(difficulty []int, profit []int, left, right int) {
	// tetrmiante
	if left >= right {
		return
	}
	pivot := quickSortHelper(difficulty, profit, left, right)
	quickSort(difficulty, profit, left, pivot-1)
	quickSort(difficulty, profit, pivot+1, right)
}

func quickSortHelper(difficulty []int, profit []int, left, right int) int {
	pivot := difficulty[right]
	cnt := left
	for ; left < right; left++ {
		if difficulty[left] < pivot {
			difficulty[cnt], difficulty[left] = difficulty[left], difficulty[cnt]
			profit[cnt], profit[left] = profit[left], profit[cnt]
			cnt++
		}
	}
	difficulty[cnt], difficulty[right] = difficulty[right], difficulty[cnt]
	profit[cnt], profit[right] = profit[right], profit[cnt]
	return cnt
}
