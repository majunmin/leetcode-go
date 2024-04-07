package leetcode_1600

type TreeNode struct {
	name    string
	deleted bool
	childs  []*TreeNode
}

func newTreeNode(name string) *TreeNode {
	return &TreeNode{
		name:    name,
		deleted: false,
		childs:  make([]*TreeNode, 0),
	}
}

// https://leetcode.cn/problems/throne-inheritance/?envType=daily-question&envId=2024-04-07
type ThroneInheritance struct {
	nodes       map[string]*TreeNode
	inheritTree *TreeNode
}

func Constructor(kingName string) ThroneInheritance {
	root := newTreeNode(kingName)
	nodes := make(map[string]*TreeNode)
	nodes[kingName] = root
	return ThroneInheritance{
		nodes:       nodes,
		inheritTree: root,
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	if node, exist := this.nodes[parentName]; exist {
		node.childs = append(node.childs, newTreeNode(childName))
	}
}

func (this *ThroneInheritance) Death(name string) {
	if node, exist := this.nodes[name]; exist {
		// add flag
		node.deleted = true
	}
}

// 节点顺序也就是树的前序遍历(出去 death节点)
func (this *ThroneInheritance) GetInheritanceOrder() []string {
	// 相当于是前序遍历了, 需要移除  已删除的节点
	//迭代的方式 输出
	var result []string
	preorder(this.inheritTree, &result)
	return result
}

func preorder(node *TreeNode, result *[]string) {
	if node == nil {
		return
	}
	if !node.deleted {
		*result = append(*result, node.name)
	}
	for _, child := range node.childs {
		preorder(child, result)
	}
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
