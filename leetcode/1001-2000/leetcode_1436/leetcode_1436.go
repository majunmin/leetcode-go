package leetcode_1436

// https://leetcode.cn/problems/destination-city/
func destCity(paths [][]string) string {
	var (
		starts = make(map[string]bool)
	)
	for _, path := range paths {
		starts[path[0]] = true
	}
	for _, path := range paths {
		if starts[path[1]] {
			continue
		}
		return path[1]
	}
	return ""
}
