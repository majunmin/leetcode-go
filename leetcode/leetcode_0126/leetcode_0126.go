package leetcode_0126

// https://leetcode.cn/problems/word-ladder-ii/description/
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return nil
	}
	if beginWord == endWord {
		return [][]string{{endWord}}
	}
	// 构建 图, 方便后面进行 dfs
	dict, minLen, finded := bfs(beginWord, endWord, wordSet)
	if !finded {
		return nil
	}
	var result [][]string
	dfs(&result, beginWord, endWord, dict, 0, minLen, []string{beginWord})
	return result

}

func dfs(result *[][]string, curWord, endWord string, dict map[string][]string, level, minLen int, path []string) {
	// terminated
	if curWord == endWord {
		dst := make([]string, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	if minLen == level {
		return
	}

	//
	if wordList, exist := dict[curWord]; exist {
		for _, next := range wordList {
			path = append(path, next)
			dfs(result, next, endWord, dict, level+1, minLen, path)
			path = path[:len(path)-1]
		}
	}
}

// @return map[stringstr][]stringstr 记录中间变化的过程
func bfs(beginWord string, endWord string, wordSet map[string]bool) (map[string][]string, int, bool) {
	//var result [][]stringstr
	queue := []string{beginWord}
	visited := make(map[string]struct{})
	dict := make(map[string][]string)
	visited[beginWord] = struct{}{}
	//visited[endWord] = struct{}{}
	minLen := 1
	found := false
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curWord := queue[i]
			for j := 0; j < len(curWord); j++ {
				for m := 0; m < 26; m++ {
					nextWord := curWord[:j] + string(rune('a'+m)) + curWord[j+1:]
					if !wordSet[nextWord] || nextWord == curWord {
						continue
					}

					if nextWord == endWord {
						found = true
					}

					if _, exist := visited[nextWord]; exist {
						continue
					}

					if len(dict[curWord]) == 0 || !containsString(dict[curWord], nextWord) {
						dict[curWord] = append(dict[curWord], nextWord)
						queue = append(queue, nextWord)
					}
				}
			}
		}
		minLen++
		if found {
			break
		}
		queue = queue[size:]
		for _, word := range queue {
			visited[word] = struct{}{}
		}
	}
	return dict, minLen, found
}

func containsString(arr []string, word string) bool {
	for _, str := range arr {
		if str == word {
			return true
		}
	}
	return false
}

// dfs solution
func dfsSolution(beginWord string, endWord string, wordList []string) [][]string {
	// param check
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	if !wordSet[endWord] {
		return nil
	}

	var result [][]string
	wordSet[beginWord] = false
	//列出所有可能的结果
	backTrace(&result, beginWord, endWord, []string{beginWord}, wordSet)
	if len(result) == 0 {
		return nil
	}
	// 找出最短的结果
	minLen := len(wordList)
	for _, path := range result {
		if len(path) < minLen {
			minLen = len(path)
		}
	}

	for i := 0; i < len(result); i++ {
		if len(result[i]) != minLen {
			result = append(result[:i], result[i+1:]...)
			i--
		}
	}

	return result
}

// wordSet: 记录已经访问过的 word
func backTrace(result *[][]string, beginWord string, endWord string, path []string, wordSet map[string]bool) {
	if beginWord == endWord {
		//copy path to result
		dst := make([]string, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	// for choice in choiceList
	for curWord, visited := range wordSet {
		if !visited {
			continue
		}

		for i := range curWord {
			for j := 0; j < 26; j++ {
				word := curWord[:i] + string(rune('a'+j)) + curWord[i+1:]
				if word != beginWord {
					continue
				}
				path = append(path, curWord)
				wordSet[curWord] = false
				backTrace(result, curWord, endWord, path, wordSet)
				path = path[:len(path)-1]
				wordSet[curWord] = true
				break
			}
		}
	}
}
