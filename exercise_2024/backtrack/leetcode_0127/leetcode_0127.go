package leetcode_0127

import "fmt"

// https://leetcode.cn/problems/word-ladder/description/
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// param check
	if beginWord == endWord {
		return 0
	}
	return doubleBFSSolution(beginWord, endWord, wordList)
}

// 这里用  list表示 queue, 用 set 来表示 单子集合，是否可以加速呢?
func doubleBFSSolution(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	// paramCheck
	if beginWord == endWord || !wordSet[endWord] {
		return 0
	}

	// process
	distance := 1
	beginQueue, endQueue := []string{beginWord}, []string{endWord}
	for len(beginQueue) > 0 && len(endQueue) > 0 {
		fmt.Println(beginQueue)
		fmt.Println(endQueue)
		llen, rlen := len(beginQueue), len(endQueue)
		// 选取长度小的  queue 作为 bfs的起点(节省计算资源)
		if llen > rlen {
			beginQueue, endQueue = endQueue, beginQueue
			llen, rlen = rlen, llen
		}
		for i := 0; i < llen; i++ {
			curWord := beginQueue[i]
			for j := 0; j < len(curWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					next := curWord[:j] + string(k) + curWord[j+1:]
					// terminate check next in endQueue or not
					if checkExist(next, endQueue) {
						return distance + 1
					}
					if wordSet[next] {
						beginQueue = append(beginQueue, next)
						delete(wordSet, next)
					}
				}
			}
		}
		beginQueue = beginQueue[llen:]
		distance++
	}
	return 0
}

func checkExist(next string, queue []string) bool {
	for _, item := range queue {
		if item == next {
			return true
		}
	}
	return false
}

func BFSSolution(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}
	delete(wordSet, beginWord)
	distance := 1
	queue := make([]string, 0)
	queue = append(queue, beginWord)
	for len(queue) > 0 {
		fmt.Println(queue)
		length := len(queue)
		for i := 0; i < length; i++ {
			preWord := queue[i]
			for j := 0; j < len(preWord); j++ {
				for k := 'a'; k <= 'z'; k++ {
					item := preWord[:j] + string(k) + preWord[j+1:]
					if item == endWord {
						return distance + 1
						// end
					}
					if wordSet[item] {
						queue = append(queue, item)
						delete(wordSet, item)
					}
				}
			}
		}
		distance++
		// terminate
		queue = queue[length:]
	}
	return -1
}

func isIntersect(wordList1, wordList2 []string) bool {
	for i := 0; i < len(wordList1); i++ {
		for j := 0; j < len(wordList2); j++ {
			if wordList1[i] == wordList2[j] {
				return true
			}
		}
	}
	return false
}
