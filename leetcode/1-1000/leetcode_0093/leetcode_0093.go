package leetcode_0093

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/restore-ip-addresses/
func restoreIpAddresses(s string) []string {
	var (
		n      = len(s)
		result = make([]string, 0)
		items  = make([]string, 0)
		dfs    func([]string, int)
	)

	dfs = func(items []string, idx int) {
		// terminate
		if len(items) == 4 && idx == n {
			result = append(result, strings.Join(items, "."))
			return
		}
		if len(items) >= 4 || idx >= n {
			return
		}

		if s[idx] == '0' {
			items = append(items, "0")
			dfs(items, idx+1)
			return
		}
		// for choice in choiceList
		for i := 1; i <= 3; i++ {
			if idx+i > n {
				continue
			}
			num, _ := strconv.Atoi(s[idx : idx+i])
			if num > 255 {
				continue
			}
			items = append(items, s[idx:idx+i])
			dfs(items, idx+i)
			items = items[:len(items)-1]
		}
	}

	dfs(items, 0)
	return result

}
