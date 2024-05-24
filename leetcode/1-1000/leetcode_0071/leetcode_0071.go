package leetcode_0071

import "strings"

// https://leetcode.cn/problems/simplify-path/?envType=study-plan-v2&envId=top-interview-150
func simplifyPath(path string) string {
	var (
		stack = make([]string, 0)
		items = strings.Split(path, "/")
	)

	for _, item := range items {
		if item == "." || item == "" {
			continue
		} else if item == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else { // item 是非空字符串
			stack = append(stack, item)
		}
	}

	if len(stack) > 0 {
		return "/"
	}

	var sb strings.Builder
	for _, item := range stack {
		sb.WriteByte('/')
		sb.WriteString(item)
	}
	return sb.String()
}

func simplifyPath2(path string) string {
	var (
		stack    = make([]string, 0)
		sBuilder strings.Builder
	)
	path = path + "/"
	for i := range path {
		if path[i] == '/' {
			if sBuilder.Len() == 0 {
				continue
			}
			item := sBuilder.String()
			sBuilder.Reset()

			if item == "." {
				continue
			} else if item == ".." {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				}
			} else {
				stack = append(stack, item)
			}
		}
		sBuilder.WriteByte(path[i])
	}

	if len(stack) == 0 {
		return "/"
	}

	// 复用 sBuilder
	for _, item := range stack {
		sBuilder.WriteByte('/')
		sBuilder.WriteString(item)
	}
	return sBuilder.String()
}
