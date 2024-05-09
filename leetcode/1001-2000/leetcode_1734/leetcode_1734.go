package leetcode_1734

func decode(encoded []int) []int {
	// param check
	if len(encoded)&1 == 1 || len(encoded) < 2 {
		panic("invalid param")
	}

	//统计 perms 所有数字的 异或 (mask)
	var (
		mask = 0
		n    = len(encoded)
		perm = make([]int, n+1)
	)
	for i := 1; i <= n; i++ {
		mask ^= i
	}
	// cal perm[0]
	var xorExclude0 int
	for i := 1; i < n; i += 2 {
		xorExclude0 ^= encoded[i]
	}
	perm[0] = mask ^ xorExclude0
	for i := 1; i <= n; i++ {
		perm[i] = perm[i-1] ^ encoded[i-1]
	}
	return perm
}
