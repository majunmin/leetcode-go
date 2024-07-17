package leetcode_3211

// https://leetcode.cn/problems/generate-binary-strings-without-adjacent-zeros/
// 1. dfs
// 2. 回溯
func validStrings(n int) []string {
	var (
		result = make([]string, 0)
	)
	// dfs
	generate(&result, n, "")
	return result
}

func generate(result *[]string, n int, cur string) {
	// terminate
	if n == 0 {
		*result = append(*result, cur)
	}
	// choose 0
	if len(cur) == 0 || cur[len(cur)-1] == '1' {
		generate(result, n-1, cur+"0")
	}
	// choose 1
	generate(result, n-1, cur+"1")
}
