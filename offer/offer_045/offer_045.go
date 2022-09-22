package offer_045

import (
	"sort"
	"strconv"
	"strings"
)

//https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/?envType=study-plan&id=lcof
func minNumber(nums []int) string {
	tempArr := make([]string, 0, len(nums))
	for i, num := range nums {
		tempArr[i] = strconv.Itoa(num)
	}
	sort.Slice(tempArr, func(i, j int) bool {
		return strings.Compare(tempArr[i], tempArr[j]) < 0
	})

	var buf strings.Builder
	for _, item := range tempArr {
		buf.WriteString(item)
	}
	return buf.String()
}
