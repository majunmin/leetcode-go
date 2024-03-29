package leetcode_0126

// 1. bfs找到最短路径  path
// 2. dfs构建 根据 bfs 搜索到的path 构建结果
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	// param check
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if len(wordSet) == 0 || !wordSet[endWord] {
		return nil
	}
	delete(wordSet, beginWord)

	var levels []map[string]bool
	found := bfs(&levels, beginWord, endWord, wordSet)
	if !found {
		return nil
	}
	// dfs 构建结果集
	var result [][]string
	dfs(&result, endWord, levels, 1, beginWord, []string{beginWord})
	return result
}

func bfs(levels *[]map[string]bool, beginWord string, endWord string, wordSet map[string]bool) bool {
	queue := []string{beginWord}

	for len(queue) > 0 {
		addToLevels(levels, queue)
		length := len(queue)
		for i := 0; i < length; i++ {
			curWord := queue[i]
			nextWords := findOneCharDiffWords(curWord, wordSet)
			for _, next := range nextWords {
				// terminate
				if next == endWord {
					*levels = append(*levels, map[string]bool{endWord: true})
					return true
				}
				queue = append(queue, next)
				delete(wordSet, next)
			}
		}
		queue = queue[length:]
	}
	return false
}

func dfs(result *[][]string, endWord string, levels []map[string]bool, nextIdx int, curWord string, path []string) {
	if nextIdx == len(levels) {
		temp := make([]string, len(path))
		copy(temp, path)
		*result = append(*result, temp)
		return
	}

	// for choice in choiceList
	for _, word := range findOneCharDiffWords(curWord, levels[nextIdx]) {
		path = append(path, word)
		dfs(result, endWord, levels, nextIdx+1, word, path)
		path = path[:len(path)-1]
	}
}

func addToLevels(levels *[]map[string]bool, queue []string) {
	set := make(map[string]bool)
	for _, word := range queue {
		set[word] = true
	}
	*levels = append(*levels, set)
}

func findOneCharDiffWords(word string, wordSet map[string]bool) []string {
	var result []string
	for i := 0; i < len(word); i++ {
		for j := 'a'; j <= 'z'; j++ {
			next := word[:i] + string(j) + word[i+1:]
			if wordSet[next] {
				result = append(result, next)
			}
		}
	}
	return result
}

// solution2
//  dfs回溯 恢复路径

func findLadders2(beginWord string, endWord string, wordList []string) [][]string {

	// param check
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if len(wordSet) == 0 || !wordSet[endWord] {
		return nil
	}
	delete(wordSet, beginWord)

	from := make(map[string]map[string]bool)
	found := bfs2(from, beginWord, endWord, wordSet)
	if !found {
		return nil
	}
	// dfs 构建结果集
	var result [][]string
	dfs2(&result, beginWord, from, endWord, []string{})
	return result
}

// 构建 图 graph
// from 记录 children -> parent 的映射关系(方便后面 dfs 进行遍历路径的构建)
func bfs2(from map[string]map[string]bool, beginWord string, endWord string, wordSet map[string]bool) bool {
	queue := []string{beginWord}
	//记录了 每个单词 记录在了哪一层里面
	steps := make(map[string]int)
	step := 0
	steps[beginWord] = step
	var flag bool

	for len(queue) > 0 {
		length := len(queue)
		step++
		for i := 0; i < length; i++ {
			curWord := queue[i]
			for j := 0; j < len(curWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					next := curWord[:j] + string(k) + curWord[j+1:]
					if _, exist := steps[next]; exist && steps[next] == step {
						from[next][curWord] = true
					}
					// terminate if
					if next == endWord {
						flag = true
					}
					if wordSet[next] {
						if from[next] == nil {
							from[next] = make(map[string]bool)
						}
						from[next][curWord] = true
						steps[next] = step
						queue = append(queue, next)
						delete(wordSet, next)
					}
				}
			}
		}
		queue = queue[length:]
		if flag {
			return true
		}
		step++
	}
	return false
}

func dfs2(result *[][]string, endWord string, from map[string]map[string]bool, curWord string, path []string) {
	// terminate
	if curWord == endWord {
		temp := make([]string, len(path))
		copy(temp, path)
		temp = append([]string{endWord}, temp...)
		*result = append(*result, temp)
		return
	}
	for parentWord := range from[curWord] {
		path = append([]string{curWord}, path...)
		dfs2(result, endWord, from, parentWord, path)
		path = path[1:]
	}
}
