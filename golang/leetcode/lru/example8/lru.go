package cache

import (
	"errors"
	"sync"
	"time"
)

// LRUCache ...
type LRUCache struct {
	mux      sync.Mutex
	size     int
	capacity int
	cacheMap sync.Map
	head     *dLinkedNode
	tail     *dLinkedNode
	ttl      time.Duration
}

type dLinkedNode struct {
	key       interface{}
	value     interface{}
	prev      *dLinkedNode
	next      *dLinkedNode
	expiredAt time.Time
}

// Constructor ...
func Constructor(capacity int, ttl time.Duration) (*LRUCache, error) {
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

	c := &LRUCache{
		mux:      sync.Mutex{},
		size:     0,
		capacity: capacity,
		cacheMap: sync.Map{},
		head:     head,
		tail:     tail,
		ttl:      ttl,
	}
	go clear(c)
	return c, nil
}

func clear(c *LRUCache) {
	for {
		c.mux.Lock()

		now := time.Now()

		var (
			firstDeleteNode *dLinkedNode
			deleteNodeList  []*dLinkedNode
		)
		for start := c.head.next; start.next != nil; start = start.next {
			if start.expiredAt.Before(now) {
				if firstDeleteNode == nil {
					firstDeleteNode = start
				}
				deleteNodeList = append(deleteNodeList, start)
			}
		}

		if firstDeleteNode != nil {
			firstDeleteNode.prev.setNext(c.tail)
		}

		for _, deleteNode := range deleteNodeList {
			deleteNode.prev = nil
			deleteNode.next = nil
			c.cacheMap.Delete(deleteNode.key)
		}

		c.mux.Unlock()
		time.Sleep(1 * time.Second)
	}
}

// Get ...
func (c *LRUCache) Get(key interface{}) interface{} {
	c.mux.Lock()
	defer c.mux.Unlock()
	if v, ok := c.cacheMap.Load(key); !ok {
		return nil
	} else {
		node := v.(*dLinkedNode)
		if node.expiredAt.Before(time.Now()) {
			node.remove()
			return nil
		}

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
			key:       key,
			value:     value,
			expiredAt: time.Now().Add(c.ttl),
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
		node.expiredAt = time.Now().Add(c.ttl)

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
