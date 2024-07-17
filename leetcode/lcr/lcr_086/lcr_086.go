package lcr_086

func partition(s string) [][]string {
	var (
		result = make([][]string, 0)
	)
	dfs(&result, s, 0, []string{})
	return result

}

func dfs(result *[][]string, s string, begin int, path []string) {
	// termiante
	if begin == len(s) {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}
	// for choice in choiceList
	for i := begin + 1; i <= len(s); i++ {
		item := s[begin:i]
		if !isPalindrome(item) {
			continue
		}
		path = append(path, item)
		dfs(result, s, i, path)
		path = path[:len(path)-1]
	}
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
