package leetcode_0763

// https://leetcode.cn/problems/partition-labels/?envType=study-plan-v2&envId=top-100-liked
func partitionLabels(s string) []int {
	bytesIdx := make(map[byte]int)
	for i := range s {
		bytesIdx[s[i]] = i
	}
	//
	var result []int
	var start, end int
	for i := 0; i < len(s); i++ {
		end = max(end, bytesIdx[s[i]])
		if end == i {
			result = append(result, i-start+1)
			start = i + 1
			continue
		}
	}
	return result
}
