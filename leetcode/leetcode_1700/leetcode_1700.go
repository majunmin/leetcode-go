package leetcode_1700

// https://leetcode.cn/problems/number-of-students-unable-to-eat-lunch/
// 简单模拟法
// 喜欢  圆形三明治的人数 可以 标记为  s0
// 喜欢  方形三明治的人数 可以 标记为  s1
// 当学生中没有喜欢  栈顶的三明治的时候, 退出 ,返回  s0 + s1
func countStudents(students []int, sandwiches []int) int {
	// param check
	if len(students) == 0 {
		return 0
	}
	s0, s1 := 0, 0
	for _, item := range students {
		s1 += item
	}
	s0 = len(students) - s1

	for _, item := range sandwiches {
		if item == 1 && s1 == 0 {
			break
		}
		if item == 0 && s0 == 0 {
			break
		}
		if item == 1 {
			s1 -= 1
		} else {
			s0 -= 1
		}
	}
	return s0 + s1
}
