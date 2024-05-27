package leetcode_0383

// https://leetcode.cn/problems/ransom-note/?envType=study-plan-v2&envId=top-interview-150
func canConstruct(ransomNote string, magazine string) bool {
	var (
		rs = make([]int, 26)
		ms = make([]int, 26)
	)
	for i := range ransomNote {
		rs[ransomNote[i]-'a']++
	}
	for i := range magazine {
		ms[magazine[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		if rs[i] > ms[i] {
			return false
		}
	}
	return true
}
