package leetcode_0898

// https://leetcode.cn/problems/bitwise-ors-of-subarrays/
func subarrayBitwiseORs(arr []int) int {
	curSet := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		curSet[arr[i]] = true
		for j := i - 1; j >= 0 && arr[i]|arr[j] != arr[j]; j-- {
			arr[j] |= arr[i]
			curSet[arr[j]] = true
		}
	}
	//
	return len(curSet)
}
