package leetcode_0278

var (
	target = 4
)

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// https://leetcode.cn/problems/first-bad-version/
func firstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := left + (right-left)>>1
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isBadVersion(version int) bool {
	return version >= target
}
