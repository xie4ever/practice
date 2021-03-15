package cache

import (
	"errors"
	"sync"
)

// LRUCache ...
type LRUCache struct {
	mux      sync.Mutex
	size     int
	capacity int
	cacheMap map[interface{}]*dLinkedNode
	head     *dLinkedNode
	tail     *dLinkedNode
}

type dLinkedNode struct {
	key   interface{}
	value interface{}
	prev  *dLinkedNode
	next  *dLinkedNode
}

// Constructor ...
func Constructor(capacity int) (*LRUCache, error) {
	if capacity <= 0 {
		return nil, errors.New("invalid capacity")
	}

	head := &dLinkedNode{
		key:   0,
		value: 0,
	}
	tail := &dLinkedNode{
		key:   0,
		value: 0,
	}
	head.next = tail
	tail.prev = head

	return &LRUCache{
		mux:      sync.Mutex{},
		size:     0,
		capacity: capacity,
		cacheMap: map[interface{}]*dLinkedNode{},
		head:     head,
		tail:     tail,
	}, nil
}

// Get ...
func (c *LRUCache) Get(key interface{}) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	if node, ok := c.cacheMap[key]; !ok {
		return nil
	} else {
		node.remove()
		node.setNext(c.head.next)
		node.setPrev(c.head)
		return node.value
	}
}

// Put ...
func (c *LRUCache) Put(key interface{}, value interface{}) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if node, ok := c.cacheMap[key]; !ok {
		deleteNode := c.tail.prev
		node := &dLinkedNode{
			key:   key,
			value: value,
		}
		c.cacheMap[key] = node

		node.setNext(c.head.next)
		node.setPrev(c.head)

		c.size++
		if c.size > c.capacity {
			c.size--
			delete(c.cacheMap, deleteNode.key)
			deleteNode.remove()
		}
		return
	} else {
		node.value = value

		node.remove()
		node.setNext(c.head.next)
		node.setPrev(c.head)
	}
}

// 移除节点，连接前后节点，前后指针设为nil
func (n *dLinkedNode) remove() {
	n.prev.next = n.next
	n.next.prev = n.prev
	n.prev = nil
	n.next = nil
}

// 绑定前节点
func (n *dLinkedNode) setPrev(prev *dLinkedNode) {
	n.prev = prev
	prev.next = n
}

// 绑定后节点
func (n *dLinkedNode) setNext(next *dLinkedNode) {
	n.next = next
	next.prev = n
}
