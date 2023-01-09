package leetcode_0493

func reversePairs(nums []int) int {

	length := len(nums)
	if length <= 1 {
		return 0
	}

	n1 := append([]int(nil), nums[:length/2]...)
	n2 := append([]int(nil), nums[length/2:]...)
	cnt := reversePairs(n1) + reversePairs(n2)
	// 统计重要反转对 n1, n2 的数量
	j := 0

	for _, v := range n1 {
		for j < len(n2) && v > 2*n2[j] {
			j++
		}
		cnt += j
	}

	// 归并 n1, n2
	p1, p2 := 0, 0
	for i := range nums {
		if p1 < len(n1) && (p2 == len(n2) || n1[p1] <= n2[p2]) {
			nums[i] = n1[p1]
			p1++
		} else {
			nums[i] = n2[p2]
			p2++
		}
	}
	return cnt

}
