package leetcode_1441

// https://leetcode.cn/problems/build-an-array-with-stack-operations/
func buildArray(target []int, n int) []string {
	// param check
	if len(target) == 0 || target[len(target)-1] > n {
		return nil
	}

	var result []string
	idx, n := 0, len(target)
	for i := 1; i <= n && idx < n; i++ {
		result = append(result, "Push")
		if i != target[idx] {
			result = append(result, "Pop")
			continue
		}
		idx++
	}
	return result
}
