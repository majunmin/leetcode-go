package leetcode_0744

// https://leetcode.cn/problems/find-smallest-letter-greater-than-target/
func nextGreatestLetter(letters []byte, target byte) byte {
	// param check
	// ...
	var (
		n           = len(letters)
		left, right = -1, n
	)
	for left+1 < right {
		mid := left + (right-left)>>1
		if letters[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	if right == n || letters[right] != target {
		return letters[0]
	}
	return letters[right]
}
