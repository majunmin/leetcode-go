package lcr_139

// https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func trainingPlan(actions []int) []int {
	//类似于快排中的一个步骤
	var (
		cnt int
	)
	for i := 0; i < len(actions); i++ {
		if actions[i]&1 == 1 {
			actions[i], actions[cnt] = actions[cnt], actions[i]
			cnt++
		}
	}
	return actions
}
