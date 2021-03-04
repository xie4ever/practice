package cache

// LRUCache ...
type LRUCache struct {
	size     int
	capacity int
	cacheMap map[int]*dLinkedNode
	head     *dLinkedNode
	tail     *dLinkedNode
}

type dLinkedNode struct {
	key   int
	value int
	prev  *dLinkedNode
	next  *dLinkedNode
}

// Constructor ...
func Constructor(capacity int) LRUCache {
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

	return LRUCache{
		size:     0,
		capacity: capacity,
		cacheMap: map[int]*dLinkedNode{},
		head:     head,
		tail:     tail,
	}
}

// Get ...
func (c *LRUCache) Get(key int) int {
	if _, ok := c.cacheMap[key]; !ok {
		return -1
	}

	node := c.cacheMap[key]

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = c.head
	node.next = c.head.next

	c.head.next.prev = node
	c.head.next = node

	return node.value
}

// Put ...
func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cacheMap[key]; !ok {
		deleteNode := c.tail.prev
		node := &dLinkedNode{
			key:   key,
			value: value,
			prev:  c.head,
			next:  c.head.next,
		}
		c.cacheMap[key] = node

		c.head.next.prev = node
		c.head.next = node

		c.size++
		if c.size > c.capacity {
			c.size--
			delete(c.cacheMap, deleteNode.key)
			c.tail.prev.prev.next = c.tail
			c.tail.prev = c.tail.prev.prev
		}
		return
	}

	node := c.cacheMap[key]
	node.value = value

	node.prev.next = node.next
	node.next.prev = node.prev

	node.prev = c.head
	node.next = c.head.next

	c.head.next.prev = node
	c.head.next = node
}
