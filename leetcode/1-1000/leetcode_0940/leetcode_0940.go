package leetcode_0940

// abc
//                  new  cnt
// a                +1    1
// ab b             +2    3
// ac abc bc c      +4    7

// https://leetcode.cn/problems/distinct-subsequences-ii/
func distinctSubseqII(s string) int {
	//parameter check
	if len(s) == 0 {
		return 0
	}
	// 用数组优化  map
	m := 1000000007
	record := make([]int, 26)
	cnt, newCnt := 0, 0
	for i := 0; i < len(s); i++ {
		b := s[i]
		newCnt = cnt + 1
		cnt = ((cnt+newCnt)%m - (record[b-'a'])%m + m) % m
		record[b-'a'] = newCnt
	}
	return cnt
}

func solution_map(s string) int {
	// 总数记录 cnt
	// 每次新添加的 item = newCnt
	// cnt = preCnt + newCnt - repeat(== 上一次该字符出现时的 newCnt) 用map记录
	m := 1000000007
	cm := make(map[byte]int)
	cnt, newCnt := 0, 0

	for i := 0; i < len(s); i++ {
		b := s[i]
		newCnt = cnt + 1
		cnt = (cnt + newCnt) % m
		// 如果 存在重复
		if _, exist := cm[b]; exist {
			cnt = (cnt - cm[b] + m) % m
		}
		cm[b] = newCnt
	}
	return cnt
}
