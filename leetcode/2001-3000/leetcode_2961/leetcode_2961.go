package leetcode_2961

// https://leetcode.cn/problems/double-modular-exponentiation/
func getGoodIndices(variables [][]int, target int) []int {
	var (
		result = make([]int, 0, len(variables))
	)
	for i, variable := range variables {
		if powMod(powMod(variable[0], variable[1], 10), variable[2], variable[3]) == target {
			result = append(result, i)
		}
	}
	return result
}

func powMod(x, y, mod int) int {
	res := 1
	for y > 0 {
		if y&1 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y >>= 1
	}
	return res
}
