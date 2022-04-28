// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-28 09:14
package leetcode_0088

import (
	"fmt"
	"testing"
)

func Test_Merge(t *testing.T) {
	nums1 := []int{2, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
