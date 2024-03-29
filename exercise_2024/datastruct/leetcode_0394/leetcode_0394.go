package leetcode_0394

type pair struct {
	prefix []byte
	repeat int
}

// https://leetcode.cn/problems/decode-string/?envType=study-plan-v2&envId=top-100-liked
func decodeString(s string) string {
	return recursiveSolution(s)
}

type pair2 struct {
	res string
	i   int // 退出递归时 ] 的索引位置
}

func recursiveSolution(s string) string {
	return dfs(s, 0).res
}

func dfs(s string, i int) pair2 {
	var (
		result string
		multi  int
	)

	for i < len(s) {
		if s[i] >= '0' && s[i] <= '9' {
			multi = multi*10 + int(s[i]-'0')
			i++
		} else if s[i] == '[' {
			// 递归
			p := dfs(s, i+1)
			for multi > 0 {
				result += p.res
				multi--
			}
			i = p.i + 1
		} else if s[i] == ']' {
			//terminate
			return pair2{res: result, i: i}
		} else if s[i] >= 'a' && s[i] <= 'z' {
			// append
			result += string(s[i])
			i++
		}
	}
	return pair2{res: result, i: i}
}

func stackSolution(s string) string {
	//  param check
	var (
		res    []byte
		repeat int
		stack  = make([]pair, 0)
	)

	for i := range s {
		if '0' <= s[i] && s[i] <= '9' {
			repeat = repeat*10 + int(s[i]-'0')
			continue
		}
		// s[i] == '[',入栈
		if s[i] == '[' {
			stack = append(stack, pair{res, repeat})
			res = []byte{}
			repeat = 0
			continue
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			res = append(res, s[i])
			continue
		}
		if s[i] == ']' {
			// cal result
			// pop stack
			item := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			tmp := item.prefix
			for j := 0; j < item.repeat; j++ {
				tmp = append(tmp, res...)
			}
			res = tmp
		}
	}
	return string(res)
}
