package lcp_0061

// https://leetcode.cn/problems/6CE719/
func temperatureTrend(temperatureA []int, temperatureB []int) int {
	var (
		n      = len(temperatureA)
		result int
		cur    int
	)
	for i := 1; i < n; i++ {
		incrA := temperatureA[i] - temperatureA[i-1]
		incrB := temperatureB[i] - temperatureB[i-1]
		if incrA == 0 && incrB == 0 || incrA*incrB > 0 {
			cur += 1
		} else {
			cur = 0
		}
		result = max(result, cur)
	}
	return result
}
