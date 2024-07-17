package leetcode_0246

// https://leetcode.ca/all/246.html
// https://leetcode.cn/problems/strobogrammatic-number/description/
var (
	itemMap = map[byte]byte{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '9',
		'9': '9',
	}
)

func isStrobogrammatic(num string) bool {
	l, r := 0, len(num)-1
	for l <= r {
		if num[r] != itemMap[num[l]] {
			return false
		}
		l++
		r--
	}
	return true
}

func isStrobogrammatic2(num string) bool {
	d := []int{0, 1, -1, -1, -1, -1, 9, -1, 8, 6}
	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		a, b := int(num[i]-'0'), int(num[j]-'0')
		if d[a] != b {
			return false
		}
	}
	return true
}
