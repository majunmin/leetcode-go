package leetcode_1095

// https://leetcode.cn/problems/find-in-mountain-array/

type MountainArray struct {
}

func (this *MountainArray) get(index int) int { return 1 }
func (this *MountainArray) length() int       { return 1 }

// https://leetcode.cn/problems/find-in-mountain-array/
func findInMountainArray(target int, mountainArr *MountainArray) int {
	topIdx := findMountainTop(mountainArr)
	res := findSortedArray(target, mountainArr, 0, topIdx)
	if res != -1 {
		return res
	}
	return findReverseSortedArray(target, mountainArr, topIdx+1, mountainArr.length()-1)
}

func findReverseSortedArray(target int, arr *MountainArray, l int, r int) int {
	l--
	r++
	for l+1 < r {
		mid := l + (r-l)>>1
		if target <= arr.get(mid) {
			l = mid
		} else {
			r = mid
		}
	}
	if l > -1 && arr.get(l) == target {
		return l
	}
	return -1
}

func findSortedArray(target int, arr *MountainArray, l int, r int) int {
	// 在有序数组 arr[l,r] 内寻找target.d
	l--
	r++
	for l+1 < r {
		mid := l + (r-l)>>1
		if arr.get(mid) <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if l > -1 && arr.get(l) == target {
		return l
	}
	return -1
}

// 寻找山脉数组峰值索引
func findMountainTop(arr *MountainArray) int {
	var (
		l, r = -1, arr.length()
	)
	for l+1 < r {
		mid := l + (r-l)>>1
		if arr.get(mid) > arr.get(mid+1) {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}
