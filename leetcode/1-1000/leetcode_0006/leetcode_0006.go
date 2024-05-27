package leetcode_0006

import "strings"

// https://leetcode.cn/problems/zigzag-conversion/?envType=study-plan-v2&envId=top-interview-150
func convert(s string, numRows int) string {
	var (
		strBuillders = make([]strings.Builder, numRows)
		idx          = 0
		flag         = 1
	)
	for i := 0; i < len(s); i++ {
		strBuillders[idx].WriteByte(s[i])
		if i == 0 || i == len(strBuillders)-1 {
			flag = -flag
		}
		idx += flag
	}
	// build result
	var sb strings.Builder
	for _, item := range strBuillders {
		sb.WriteString(item.String())
	}
	return sb.String()
}
