package leetcode_0438

import "reflect"

func findAnagrams(s string, p string) []int {
	// param check
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return nil
	}
	//skip
	pArr := make([]int, 26)
	for i := range p {
		pArr[p[i]-'a']++
	}
	var result []int
	l, r := 0, 0
	//process
	sArr := make([]int, 26)
	for r < sLen {
		sArr[s[r]-'a']++
		// shrink window
		if r-l+1 == pLen {
			if reflect.DeepEqual(sArr, pArr) {
				result = append(result, l)
			}
			sArr[s[l]-'a']--
			l++
		}
		r++
	}
	return result

}

// golang 判断数组是否相同
//reflect.DeepEqual()
