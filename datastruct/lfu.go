package datastruct

type LFUCache struct {
	capacity int
	cache    map[string]*Node
	fList    []*DoubleLinkList // 维护访问频次
	size     int
}

func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		cache:    make(map[string]*Node, capacity),
		fList:    make([]*DoubleLinkList, 0),
		size:     0,
	}
}

func (c *LFUCache) Get(key string) string {
	if c.size == 0 {
		return ""
	}

	node, exist := c.cache[key]
	if !exist {
		return ""
	}
	c.removeNode(node)
	node.frequency++
	if len(c.fList) < node.frequency {
		c.fList = append(c.fList, NewDoubleLinkList())
	}
	c.addNode(node)
	return node.value
}

func (c *LFUCache) addNode(node *Node) {
	c.fList[node.frequency].AddFirst(node)
	c.size++
}

func (c *LFUCache) removeNode(node *Node) {
	c.fList[node.frequency].removeNode(node)
	c.size--
}

func (c *LFUCache) Put(key, value string) {
	node, exist := c.cache[key]
	if !exist {
		node := &Node{
			frequency: 0,
			key:       key,
			value:     value,
		}
		if c.size == c.capacity {
			node := c.removeOldest()
			delete(c.cache, node.key)
		}
		c.cache[key] = node
		c.fList[node.frequency].AddFirst(node)
	}
	c.removeNode(node)
	node.frequency++
	if len(c.fList) == node.frequency {
		c.fList = append(c.fList, NewDoubleLinkList())
	}
	c.addNode(node)
}

func (c *LFUCache) removeOldest() *Node {
	for _, dl := range c.fList {
		if dl.size == 0 {
			continue
		}
		return dl.removeOldest()
	}
	return nil
}
