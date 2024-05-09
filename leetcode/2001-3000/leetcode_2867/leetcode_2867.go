package leetcode_2867

const N = 10001

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

// https://leetcode.cn/problems/count-valid-paths-in-a-tree/
func countPaths(n int, edges [][]int) int64 {
	adj := make([][]int, n)
	// 构建无向图
	for _, item := range edges {
		adj[item[0]] = append(adj[item[0]], item[1])
		adj[item[1]] = append(adj[item[1]], item[])
	}
	// dfs
	var (
		seen  = make([]int, 0)
		count = make([]int64, n+1)
		res   int64
	)
	for i := 1; i <= n; i++ {
		if !isPrime(i) {
			continue
		}
		var cur int64
		for _, j := range adj[i] {
			if isPrime(j) {
				continue
			}
			// process
			if count[j] == 0 {
				seen = seen[:0]
				dfs(adj, &seen, j, 0)
				cnt := int64(len(seen))
				for _, k := range seen {
					count[k] = cnt
				}
			}
			res += count[j] * cur
			cur += count[j]
		}
		res += cur
	}
	return res

}

func dfs(adj [][]int, seen *[]int, i int, pre int) {
	*seen = append(*seen, i)
	for _, j := range adj[i] {
		if j != pre && !isPrime(j) {
			dfs(adj, seen, j, i)
		}
	}
}
