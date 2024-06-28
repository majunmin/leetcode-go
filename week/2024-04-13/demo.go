package _024_04_13

func scoreOfString(s string) int {
	var result int
	for i := 1; i < len(s); i++ {
		result += int(s[i] - s[i-1])
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
