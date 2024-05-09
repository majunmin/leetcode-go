package leetcode_1442

// a = preX[i] ^ preX[j]
// b = preX[j] ^ preX[k+1]
// condition : a == b ==> a^b == 0  ==> preX[i] == preX[k+1]

// https://leetcode.cn/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/
func countTriplets(arr []int) int {
	n := len(arr)
	preX := make([]int, n+1)
	for i := 0; i < n; i++ {
		preX[i+1] = preX[i] ^ arr[i]
	}
	var cnt int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if preX[i] == preX[k+1] {
					cnt++
				}
			}
		}
	}
	return cnt
}

// 二重循环
func countTriplets2(arr []int) int {
	n := len(arr)
	preX := make([]int, n+1)
	for i := 0; i < n; i++ {
		preX[i+1] = preX[i] ^ arr[i]
	}
	var cnt int
	for i := 0; i < n-1; i++ {
		for k := i + 1; k < n; k++ {
			if preX[i] == preX[k+1] {
				cnt += k + 1 - i
			}
		}
	}
	return cnt
}

// 前缀和 + map优化题型
func countTriplets3_1(arr []int) int {
	n := len(arr)
	preX := make([]int, n+1)
	for i := 0; i < n; i++ {
		preX[i+1] = preX[i] ^ arr[i]
	}
	var (
		cnts   = make(map[int]int)
		total  = make(map[int]int)
		result int
	)
	for k := 0; k < n; k++ {
		if _, exist := cnts[preX[k+1]]; exist {
			result += cnts[preX[k+1]*k-total[preX[k+1]]]
		}
		cnts[preX[k]] += 1
		total[preX[k]] += k
	}

	return result
}
