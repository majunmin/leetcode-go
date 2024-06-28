package leetcode_100345

func minimumOperations(nums []int) int {
	var (
		result int
	)
	for _, num := range nums {
		mod := num % 3
		result += min(mod, 3-mod)
	}
	return result
}
