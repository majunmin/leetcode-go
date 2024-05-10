package leetcode_2960

// https://leetcode.cn/problems/count-tested-devices-after-test-operations/
func countTestedDevices(batteryPercentages []int) int {
	var (
		result int
		decr   int
	)
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i]-decr > 0 {
			decr++
			result++
		}
	}
	return result
}
