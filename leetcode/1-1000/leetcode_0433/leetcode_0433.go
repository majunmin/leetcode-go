package leetcode_0433

import "math"

// https://leetcode.cn/problems/minimum-genetic-mutation/?envType=study-plan-v2&envId=top-interview-150
func minMutation(startGene string, endGene string, bank []string) int {
	var (
		bankSet = make(map[string]bool)
		seen    = make(map[string]bool)
	)
	for _, item := range bank {
		bankSet[item] = true
	}
	bankSet[startGene] = true
	if !bankSet[endGene] {
		return -1
	}
	seen[startGene] = true
	res := backtrace(bankSet, startGene, endGene, seen)
	if res == math.MaxInt {
		return -1
	}
	return res
}

func backtrace(bankSet map[string]bool, curGene string, targetGene string, seen map[string]bool) int {
	// terminate
	if curGene == targetGene {
		return 0
	}
	res := math.MaxInt
	for next := range bankSet {
		if !seen[next] && isChangeOne(next, curGene) {
			seen[next] = true
			pathCnt := backtrace(bankSet, next, targetGene, seen)
			if pathCnt != math.MaxInt {
				res = min(res, pathCnt+1)
			}
			// revert
			seen[next] = false
		}
	}
	return res
}

func isChangeOne(next string, gene string) bool {
	var cnt int
	for i := 0; i < len(gene); i++ {
		if next[i] != gene[i] {
			cnt++
		}
	}
	return cnt == 1
}
