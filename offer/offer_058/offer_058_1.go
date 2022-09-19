package offer_058

import (
	"strings"
)

func reverseWords(s string) string {
	strArr := strings.Fields(strings.TrimSpace(s))
	i, j := 0, len(strArr)-1
	for i < j {
		strArr[i], strArr[j] = strArr[j], strArr[i]
		i++
		j--
	}

	return strings.Join(strArr, " ")
}
