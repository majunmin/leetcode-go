package leetcode_1487

import "strconv"

// https://leetcode.cn/problems/making-file-names-unique/
func getFolderNames(names []string) []string {
	var (
		result = make([]string, 0)
		index  = make(map[string]int)
	)
	for _, name := range names {
		if _, exist := index[name]; exist {
			index[name] = 1
			result = append(result, name)
			continue
		}
		// 出现重复名字了.
		k := index[name]
		for index[addSuffix(name, k)] > 0 {
			k++
		}
		newName := addSuffix(name, k)
		index[newName] = k + 1
		result = append(result, newName)
	}
	return result
}

func addSuffix(name string, k int) string {
	return name + "(" + strconv.Itoa(k) + ")"
}

// 为什么会超时呢?
func solution1(names []string) []string {
	// 1. 用一个map 记录已经出现过的文件名.
	var (
		result  []string
		nameSet = make(map[string]bool) // 记录文件名
	)
	for _, name := range names {
		if !nameSet[name] {
			result = append(result, name)
			nameSet[name] = true
			continue
		}
		result = append(result, createFile(name, 1, nameSet))
	}
	return result
}

func createFile(name string, cnt int, nameSet map[string]bool) string {
	newName := name + "(" + strconv.Itoa(cnt) + ")"
	if !nameSet[newName] {
		nameSet[newName] = true
		return newName
	}
	return createFile(name, cnt+1, nameSet)
}
