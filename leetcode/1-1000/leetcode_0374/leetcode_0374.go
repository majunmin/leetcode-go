package leetcode_0374

var (
	target = 6
)

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
// https://leetcode.cn/problems/guess-number-higher-or-lower/
func guessNumber(n int) int {
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if guess(mid) >= 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func guess(num int) int {
	if num < target {
		return -1
	} else if num > target {
		return 1
	} else {

	}
}
