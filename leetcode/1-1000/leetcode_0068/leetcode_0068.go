package leetcode_0068

import "strings"

// https://leetcode.cn/problems/text-justification/description/?envType=study-plan-v2&envId=top-interview-150
func fullJustify(words []string, maxWidth int) []string {
	var (
		result []string
		tmp    []string
	)
	for _, word := range words {
		tmp = append(tmp, word)
		if !canContain(tmp, maxWidth) {
			// 构造 字符串， 拼接到result 中
			result = append(result, buildWidthStr(tmp[:len(tmp)-1], maxWidth))
			tmp = tmp[len(tmp)-1:]
		}
	}
	// 3. 文本最后一行的处理
	word := strings.Join(tmp, " ")
	result = append(result, word+makeSpace(maxWidth-len(word)))
	tmp = tmp[:0]
	return result
}

func buildWidthStr(strs []string, width int) string {
	var (
		n = len(strs)
	)
	if n == 1 {
		return strs[0] + makeSpace(width-len(strs[0]))
	}
	var wordTotalLen int
	for _, item := range strs {
		wordTotalLen += len(item)
	}
	// 要填充的空格数
	spaces := width - wordTotalLen

	// 平均每个位置要添加空格数:
	// width / (len(strs)-1)
	// width % (len(strs))
	prev := spaces % (n - 1)
	split := makeSpace(spaces / (n - 1))
	var sb strings.Builder
	for i, item := range strs {
		sb.WriteString(item)
		if i < n-1 {
			sb.WriteString(split)
			if i < prev {
				sb.WriteByte(' ')
			}
		}
	}
	return sb.String()
}

func makeSpace(space int) string {
	var s strings.Builder
	for i := 0; i < space; i++ {
		s.WriteByte(' ')
	}
	return s.String()
}

func canContain(tmp []string, maxWidth int) bool {
	var length int
	for _, word := range tmp {
		length += len(word)
	}
	return (length + len(tmp) - 1) <= maxWidth
}
