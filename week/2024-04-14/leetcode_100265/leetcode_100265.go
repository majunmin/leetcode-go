package leetcode_100265

const N = 101

var (
	is_prime = [N]bool{}
)

// 埃氏筛
func init() {
	for i := 2; i < N; i++ {
		is_prime[i] = true
	}
	for i := 2; i*i < N; i++ {
		if is_prime[i] {
			for j := i * i; j < N; j += i {
				is_prime[j] = false
			}
		}
	}
}

func isPrime(num int) bool {
	return is_prime[num]
}

// https://leetcode.cn/contest/weekly-contest-393/problems/maximum-prime-difference/
func maximumPrimeDifference(nums []int) int {
	l, r := 0, len(nums)-1
	for ; l < len(nums); l++ {
		if isPrime(nums[l]) {
			break
		}
	}
	//
	if l == len(nums) {
		return -1
	}
	for ; r >= l; r-- {
		if isPrime(nums[r]) {
			break
		}
	}
	return r - l
}
