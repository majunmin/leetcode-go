package leetcode_0017

var (
	bytesMap = map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
)

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/?envType=study-plan-v2&envId=top-100-liked
func letterCombinations(digits string) []string {
	// param check
	if len(digits) == 0 {
		return nil
	}
	// digits 仅包含  2-9
	var result []string
	backtrace(&result, digits, []byte{}, 0)
	return result
}

func backtrace(result *[]string, digits string, path []byte, idx int) {
	// termiante
	if len(path) == len(digits) {
		*result = append(*result, string(path))
		return
	}

	// for choice in choiceList
	for _, b := range bytesMap[digits[idx]] {
		path = append(path, b)
		backtrace(result, digits, path, idx+1)
		path = path[:len(path)-1]
	}
}
