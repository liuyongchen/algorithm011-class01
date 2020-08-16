package Week_08

type LinkNode struct {
	key, value int
	pre, next  *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	head, tail *LinkNode
	capacity   int
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.next = head
	return LRUCache{
		m:        make(map[int]*LinkNode),
		head:     head,
		tail:     tail,
		capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	m := this.m
	if node, ok := m[key]; ok {
		this.MoveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	m := this.m
	if node, ok := m[key]; ok {
		this.MoveToHead(node)
		node.value = value
	} else {
		n := &LinkNode{key, value, nil, nil}
		if len(m) >= this.capacity {
			delNode := this.tail.pre
			delete(m, delNode.key)
			this.Remove(delNode)
		}
		m[key] = n
		this.AddNode(n)

	}
}

func (this *LRUCache) AddNode(node *LinkNode) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func (this *LRUCache) Remove(node *LinkNode) {
	node.next.pre = node.pre
	node.pre.next = node.next
	node.next = nil
	node.pre = nil
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.Remove(node)
	this.AddNode(node)
}
