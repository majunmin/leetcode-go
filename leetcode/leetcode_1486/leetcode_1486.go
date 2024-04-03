package leetcode_1486

// https://leetcode.cn/problems/xor-operation-in-an-array/
// 模拟解法, 可能会溢出
func xorOperation(n int, start int) int {
	var result int
	for i := 0; i < n; i++ {
		result ^= start + 2*i
	}
	return result
}

// 数学解法 1 ^ 2 ^ 3 = 0
// 4i 异或 (4i+1) 异或 (4i+2) 异或 (4i+3)=0
// 最后一位是: n & start & 1
func xorOperation2(n int, start int) int {
	// 整体 /2
	s := start >> 1
	prefix := calc(s-1) ^ calc(s+n-1)
	// 利用奇偶性得出{最低一位}结果
	last := n & start & 1
	return prefix<<1 | last // prefix <<1 + last

}

func calc(x int) int {
	switch x % 4 {
	case 0:
		return x
	case 1:
		return 1
	case 2:
		return x + 1
	default:
		return 0
	}
}
