package main

type LRUCache struct {
	size       int
	mp         map[int]*Node
	head, tail *Node
}

type Node struct {
	key, val  int
	pre, next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.pre = head
	return LRUCache{
		size: capacity,
		mp:   map[int]*Node{},
		head: head,
		tail: tail,
	}
}

func (this *LRUCache) deleteNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) addToHead(node *Node) {
	node.next = this.head.next
	node.pre = this.head
	this.head.next = node
	node.next.pre = node
}

func (this *LRUCache) moveToHead(node *Node) {

	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.mp[key]; ok {
		this.moveToHead(v)
		return v.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if v, ok := this.mp[key]; ok {
		v.val = value
		this.moveToHead(v)
	} else {
		if this.size == len(this.mp) {
			delKey := this.tail.pre.key
			this.deleteNode(this.tail.pre)
			delete(this.mp, delKey)
		}
		// 新增节点
		addNode := &Node{key: key, val: value}
		this.addToHead(addNode)
		this.mp[key] = addNode
	}
}
