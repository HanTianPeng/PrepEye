package main

type Node struct {
	Key    int
	Val    int
	Pre    *Node
	Next   *Node
	Random *Node
}
type LRUCache struct {
	capacity  int
	dummy     *Node
	keyToNode map[int]*Node
}

func Constructor(capacity int) LRUCache {
	dummy := &Node{}
	dummy.Pre = dummy
	dummy.Next = dummy
	return LRUCache{
		capacity:  capacity,
		dummy:     dummy,
		keyToNode: make(map[int]*Node, capacity),
	}
}

// 删除一个节点
func (c *LRUCache) remove(x *Node) {
	x.Pre.Next = x.Next
	x.Next.Pre = x.Pre
}

// 头部添加一个节点
func (c *LRUCache) putHead(x *Node) {
	x.Pre = c.dummy
	x.Next = c.dummy.Next
	x.Pre.Next = x
	x.Next.Pre = x
}

// 获取key对应的节点;同时将该节点移动到头节点
func (c *LRUCache) getNode(key int) *Node {
	node := c.keyToNode[key]
	if node == nil {
		return nil
	}
	c.remove(node)
	c.putHead(node)
	return node
}

func (c *LRUCache) Get(key int) int {
	node := c.getNode(key)
	if node == nil {
		return -1
	}
	return node.Val
}

func (c *LRUCache) Put(key int, value int) {
	node := c.getNode(key)
	if node != nil {
		node.Val = value
		return
	}
	node = &Node{Key: key, Val: value}
	c.keyToNode[key] = node
	c.putHead(node)
	if len(c.keyToNode) > c.capacity {
		backNode := c.dummy.Pre
		delete(c.keyToNode, backNode.Key)
		c.remove(backNode)
	}
}
