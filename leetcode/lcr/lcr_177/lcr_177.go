package lcr_177

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func sockCollocation(sockets []int) []int {
	var ret int // a^b
	for _, soc := range sockets {
		ret ^= soc
	}
	// 找到div中 最低位的1 (任意一个 1 都可以)
	//div := 1
	//for div&ret == 0 {
	//	div <<= 1
	//}
	// lowbit
	div := ret & (-ret)
	// 分组
	var a, b int
	for _, soc := range sockets {
		if soc^div == 0 {
			a ^= soc
		} else {
			b ^= soc
		}
	}
	return []int{a, b}
}
