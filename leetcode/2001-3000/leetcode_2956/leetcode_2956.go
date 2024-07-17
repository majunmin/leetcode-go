package leetcode_2956

// https://leetcode.cn/problems/find-common-elements-between-two-arrays/
func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var (
		num1Cnts   = make(map[int]int)
		num2Cnts   = make(map[int]int)
		res1, res2 int
	)
	for _, num := range nums1 {
		num1Cnts[num]++
	}
	for _, num := range nums2 {
		num2Cnts[num]++
	}
	for num := range num1Cnts {
		res1 += num2Cnts[num]
	}
	for num := range num2Cnts {
		res2 += num1Cnts[num]
	}
	return []int{res2, res1}
}
