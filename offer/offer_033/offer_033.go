package offer_033

// https://leetcode.cn/problems/er-cha-sou-suo-shu-de-hou-xu-bian-li-xu-lie-lcof/?envType=study-plan&id=lcof
// 二叉搜索树
func verifyPostorder(postorder []int) bool {
	return verifyPostorderHelper(postorder, 0, len(postorder)-1)
}

func verifyPostorderHelper(postorder []int, left int, right int) bool {
	if left >= right {
		return true
	}

	//
	rootVal := postorder[right]
	// judge
	mid := left
	for postorder[mid] < rootVal {
		mid++
	}
	for i := mid; i < right; i++ {
		if postorder[i] < rootVal {
			return false
		}
	}

	return verifyPostorderHelper(postorder, left, mid-1) && verifyPostorderHelper(postorder, mid, right-1)

}
