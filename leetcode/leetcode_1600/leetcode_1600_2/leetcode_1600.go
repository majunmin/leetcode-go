package leetcode_1600_2

// https://leetcode.cn/problems/throne-inheritance/?envType=daily-question&envId=2024-04-07
type ThroneInheritance struct {
	root     string
	deaths   map[string]bool
	inherite map[string][]string
}

func Constructor(kingName string) ThroneInheritance {
	inherite := make(map[string][]string)
	inherite[kingName] = make([]string, 0)
	return ThroneInheritance{
		root:     kingName,
		deaths:   make(map[string]bool),
		inherite: inherite,
	}

}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.inherite[parentName] = append(this.inherite[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.deaths[name] = true
}

// 节点顺序也就是树的前序遍历(出去 death节点)
func (this *ThroneInheritance) GetInheritanceOrder() []string {
	// 相当于是前序遍历了, 需要移除  已删除的节点
	//迭代的方式 输出
	var result []string
	this.preorder(this.root, &result)
	return result
}

func (this *ThroneInheritance) preorder(node string, res *[]string) {
	if !this.deaths[node] {
		*res = append(*res, node)
	}
	for _, child := range this.inherite[node] {
		this.preorder(child, res)
	}
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
