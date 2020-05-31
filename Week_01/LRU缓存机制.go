package main

import "fmt"

// https://leetcode-cn.com/problems/lru-cache/submissions/

//第一遍: 2020-05-24 20:37
type LRUCacheNode struct {
	key  int
	val  int
	next *LRUCacheNode
	prev *LRUCacheNode
}

type LRUCache struct {
	capcaity int
	len      int
	head     *LRUCacheNode
	tail     *LRUCacheNode
	data     map[int]*LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	lc := LRUCache{
		capcaity: capacity,
		len:      0,
		head:     &LRUCacheNode{key: -1, val: -1},
		tail:     &LRUCacheNode{key: -1, val: -2},
		data:     make(map[int]*LRUCacheNode),
	}
	lc.head.next = lc.tail
	lc.tail.prev = lc.head
	return lc
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.data[key]; !ok {
		return -1
	} else {
		this.deleteNode(node)
		this.moveToHead(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.data[key]; !ok {
		if this.isFull() {
			this.removeTailNode()
			this.len--
		}
		node = &LRUCacheNode{
			key: key,
			val: value,
		}
		this.moveToHead(node)
		this.data[key] = node
		this.len++
	} else {
		node.val = value
		this.deleteNode(node)
		this.moveToHead(node)
	}
}

func (this *LRUCache) isFull() bool {
	if this.len == this.capcaity {
		return true
	}
	return false
}

func (this *LRUCache) moveToHead(node *LRUCacheNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeTailNode() {
	delete(this.data, this.tail.prev.key)
	this.deleteNode(this.tail.prev)
}

func (this *LRUCache) deleteNode(node *LRUCacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

type testData struct {
	op   string
	data []int
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	var (
		ConstructorOp = "Constructor"
		GetOp         = "Get"
		PutOp         = "Put"
	)

	testCase := [][]testData{
		{
			{ConstructorOp, []int{3}},
			{PutOp, []int{1, 1}},
			{PutOp, []int{2, 2}},
			{GetOp, []int{1}},
			{PutOp, []int{2, 2}},
			{GetOp, []int{2}},
			{PutOp, []int{3, 3}},
			{GetOp, []int{3}},
			{GetOp, []int{2}},
		},
	}

	for _, c := range testCase {
		var obj LRUCache
		for _, d := range c {
			switch d.op {
			case ConstructorOp:
				obj = Constructor(d.data[0])
			case GetOp:
				r := obj.Get(d.data[0])
				fmt.Println(r)
			case PutOp:
				obj.Put(d.data[0], d.data[1])
			}
		}
	}
}
