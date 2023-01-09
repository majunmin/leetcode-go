package leetcode_0017

var (
	letterMap = map[byte][]byte{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}
)

// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return nil
	}
	var result []string
	backTrace(&result, l, []byte(digits), []byte{}, 0)
	return result
}

func backTrace(result *[]string, l int, digits []byte, path []byte, idx int) {
	if idx == l {
		dst := make([]byte, l)
		copy(dst, path)
		*result = append(*result, string(dst))
		return
	}

	for _, letter := range letterMap[digits[idx]] {
		path = append(path, letter)
		backTrace(result, l, digits, path, idx+1)
		path = path[:len(path)-1]
	}
}

func letterCombinations2(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letterMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string

	var recur func(path string, idx int)
	recur = func(path string, idx int) {
		if idx >= len(digits) {
			result = append(result, path)
			return
		}

		s := letterMap[digits[idx]]
		for i := 0; i < len(s); i++ {
			recur(path+string(s[i]), idx+1)
		}
	}

	recur("", 0)
	return result
}
