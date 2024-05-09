// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-20 23:59
package leetcode_0717

func isOneBitCharacter(bits []int) bool {
	i := 0
	for i < len(bits) {
		// terminate
		if i == len(bits)-1 {
			return true
		}
		if bits[i] == 1 {
			i += 2
		} else {
			i++
		}
	}
	return false
}

// isOneBitChacter2
func isOneBitCharacter2(bits []int) bool {
	n := len(bits)
	i := n - 2
	for i >= 0 {
		if bits[i] == 0 {
			break
		}
		i--
	}
	// n- i 是偶数
	return (n-i)&1 == 0
}
