package leetcode_0771

// https://leetcode.cn/problems/jewels-and-stones/
func numJewelsInStones(jewels string, stones string) int {
	//param check
	if len(jewels) == 0 || len(stones) == 0 {
		return 0
	}
	// process
	jewelsMap := make(map[byte]struct{}, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jewelsMap[jewels[i]] = struct{}{}
	}

	var cnt int
	for i := range stones {
		if _, exist := jewelsMap[stones[i]]; exist {
			cnt++
		}
	}
	return cnt
}
