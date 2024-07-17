package lcr_170

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/description/
// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func reversePairs(record []int) int {
	// param check
	if len(record) < 2 {
		return 0
	}
	return mergeSort(record, 0, len(record)-1)
}

func mergeSort(record []int, l, r int) int {
	if l >= r {
		return 0
	}
	mid := l + (r-l)>>1
	// 分
	res := mergeSort(record, l, mid) + mergeSort(record, mid+1, r)
	// 合
	tmp := make([]int, r-l+1)
	for i := l; i < r; i++ {
		tmp[i] = record[i]
	}
	i, j := l, mid+1
	for k := l; k <= r; k++ {
		if k == mid+1 { // 左边的已经遍历完
			record[k] = tmp[j]
			j++
		} else if k == r+1 || tmp[i] <= tmp[j] {
			record[k] = tmp[i]
			i++
		} else { // tmp[i] > tmp[j]
			record[k] = tmp[j]
			j++
			// 左边有 mid + 1 - i 个数 > record[j]
			res += mid + 1 - i // 统计逆序对
		}
	}
	return res
}
