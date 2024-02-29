package leetcode_0204

// 埃氏筛：https://leetcode.cn/problems/count-primes/solutions/1/kuai-lai-miao-dong-shai-zhi-shu-by-sweetiee/
func countPrimes(n int) int {
	isprime := make([]bool, n)
	for i := 0; i < n; i++ {
		isprime[i] = true
	}
	var cnt int
	for i := 2; i < n; i++ {
		if isprime[i] {
			cnt++
			for j := 2 * i; j <= n; j++ {
				isprime[j] = false
			}
		}
	}
	return cnt
}

// 普通的解法， 超时
func solution1(n int) int {
	var cnt int
	for i := 1; i < n; i++ {
		if isPrime(i) {
			cnt++
		}
	}
	return cnt
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if (n/i)*i == n {
			return true
		}
	}
	return false
}
