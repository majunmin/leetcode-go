package leetcode_100256

// https://leetcode.cn/contest/weekly-contest-393/problems/latest-time-you-can-obtain-after-replacing-characters/
func findLatestTime(s string) string {
	// if else
	var (
		result = make([]byte, 5)
	)
	for i := range s {
		result[i] = s[i]
	}

	// param check
	if s[0] == '?' && s[1] == '?' {
		result[0] = '1'
		result[1] = '1'
	}
	if s[0] == '?' {
		result[0] = '0'
		if result[1] < '2' {
			result[0] = '1'
		}
	}
	if s[1] == '?' {
		result[1] = '9'
		if result[0] == '1' {
			result[1] = '1'
		}
	}
	result[2] = ':'
	if s[3] == '?' {
		result[3] = '5'
	}
	if s[4] == '?' {
		result[4] = '9'
	}
	return string(result)
}
