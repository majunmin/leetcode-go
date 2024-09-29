package main

import (
	"fmt"
	"strings"
)

package main

import (
"fmt"
"strconv"
"strings"
)

func decompress(s string) string {
	var strStack []string
	var countStack []int

	var currentStr strings.Builder
	n := len(s)
	i := 0

	for i < n {
		ch := s[i]

		if isDigit(ch) {
			// 处理数字，可能是多位数
			num := 0
			for i < n && isDigit(s[i]) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			countStack = append(countStack, num) // 将数字（重复次数）压入栈
		} else if ch == '{' {
			// 遇到 '{'，说明是新的一段，先把当前字符串入栈保存
			strStack = append(strStack, currentStr.String())
			currentStr.Reset() // 开始新的嵌套部分
			i++
		} else if ch == '}' {
			// 遇到 '}'，处理最近一个嵌套
			tempStr := currentStr.String()
			currentStr.Reset()
			lastStr := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1] // 出栈上一个字符串

			// 取出重复次数
			repeatTimes := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			// 重复字符串并附加到之前的字符串
			currentStr.WriteString(lastStr)
			for j := 0; j < repeatTimes; j++ {
				currentStr.WriteString(tempStr)
			}
			i++
		} else {
			// 处理普通字符，直接加到当前的字符串
			currentStr.WriteByte(ch)
			i++
		}
	}

	return currentStr.String() // 最终结果
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func main() {
	compressedStr := "{A3B1{C}3}3"
	decompressedStr := decompress(compressedStr)
	fmt.Println(decompressedStr)
}
