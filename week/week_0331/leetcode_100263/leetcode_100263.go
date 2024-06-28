package leetcode_100263

func sumOfTheDigitsOfHarshadNumber(x int) int {
	// param check
	if x <= 0 {
		return -1
	}
	var sum int
	num := x
	for num > 0 {
		sum += num % 10
		num /= 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
