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
	cacheMap sync.Map
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
		size:     0,
		capacity: capacity,
		cacheMap: sync.Map{},
		head:     head,
		tail:     tail,
	}, nil
}

// Get ...
func (c *LRUCache) Get(key interface{}) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	if v, ok := c.cacheMap.Load(key); !ok {
		return nil
	} else {
		node := v.(*dLinkedNode)
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
	if v, ok := c.cacheMap.Load(key); !ok {
		deleteNode := c.tail.prev
		node := &dLinkedNode{
			key:   key,
			value: value,
		}
		c.cacheMap.Store(key, node)

		node.setNext(c.head.next)
		node.setPrev(c.head)

		c.size++
		if c.size > c.capacity {
			c.size--
			c.cacheMap.Delete(deleteNode.key)
			deleteNode.remove()
		}
		return
	} else {
		node := v.(*dLinkedNode)
		node.value = value

		node.remove()
		node.setNext(c.head.next)
		node.setPrev(c.head)
	}
}

// 移除节点，连接前后节点
func (n *dLinkedNode) remove() {
	n.prev.next = n.next
	n.next.prev = n.prev
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
