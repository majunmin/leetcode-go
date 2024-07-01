package lcr_0135

import "fmt"

// 打印大数
func countNumbers(cnt int) {
	//对于大数问题.
	if cnt < 0 {
		return
	}
	printToMaxDigits([]byte{}, 0, cnt)
}

func printToMaxDigits(path []byte, cur int, n int) {
	if len(path) == n {
		printNum(path)
		return
	}
	for c := '0'; c <= '9'; c++ {
		// 0 - 0 - 0
		// 0 - 0 - 1
		path = append(path, byte(c))
		printToMaxDigits(path, cur+1, n)
		path = path[:len(path)-1]
	}
}

func printNum(path []byte) {
	skip := true
	for _, item := range path {
		if item != '0' {
			skip = false
		}
		if !skip {
			fmt.Printf("%c", item)
		}
	}
	fmt.Println()
}
