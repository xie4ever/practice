package cache

// LFUCache ...
type LFUCache struct {
	size     int
	capacity int
	maxFreq  int
	freqMap  map[int]*dLinkedNode
	valMap   map[int]*dLinkedNode
}

type dLinkedNode struct {
	key   int
	value int
	freq  int
	prev  *dLinkedNode
	next  *dLinkedNode
}

// Constructor ...
func Constructor(capacity int) LFUCache {
	return LFUCache{
		size:     0,
		capacity: capacity,
		freqMap:  make(map[int]*dLinkedNode),
		valMap:   make(map[int]*dLinkedNode),
	}
}

// Get ...
func (c *LFUCache) Get(key int) int {
	if c.capacity <= 0 {
		return -1
	}
	if node := c.valMap[key]; node == nil {
		return -1
	}

	node := c.valMap[key]

	// 已经是头节点
	if node.prev == nil {
		if node.next != nil {
			// 处理下一个节点
			node.next.prev = nil
			c.freqMap[node.freq] = node.next
		} else {
			// 没有下一个节点直接置为空
			c.freqMap[node.freq] = nil
		}

		node.freq++
		if node.freq > c.maxFreq {
			c.maxFreq = node.freq
		}

		if oldHead := c.freqMap[node.freq]; oldHead == nil {
			// 下一频率没有头节点
			node.next = nil
			node.prev = nil
			c.freqMap[node.freq] = node
		} else {
			node.next = oldHead
			oldHead.prev = node
			c.freqMap[node.freq] = node
		}

		return node.value
	}

	// 不是头结点
	node.prev.next = node.next
	if node.next != nil {
		node.next.prev = node.prev
	}

	node.freq++
	if node.freq > c.maxFreq {
		c.maxFreq = node.freq
	}

	if oldHead := c.freqMap[node.freq]; oldHead == nil {
		// 下一频率没有头节点
		node.next = nil
		node.prev = nil
		c.freqMap[node.freq] = node
	} else {
		node.next = oldHead
		oldHead.prev = node
		c.freqMap[node.freq] = node
	}

	return node.value
}

// Put ...
func (c *LFUCache) Put(key int, value int) {
	if c.capacity <= 0 {
		return
	}

	if node := c.valMap[key]; node == nil {
		c.size++
		if c.size > c.capacity {
			// 先干掉访问频率最少的那个节点
			c.size--
			for i := 1; i <= c.maxFreq; i++ {
				if node := c.freqMap[i]; node == nil {
					continue
				} else {
					// 找到最低的频率，找到最后的节点
					for node.next != nil {
						node = node.next
					}

					if node.prev == nil {
						// 如果是头节点，直接置为nil
						c.freqMap[node.freq] = nil
					} else {
						// 直接删除节点
						node.prev.next = nil
					}

					c.valMap[node.key] = nil
					break
				}
			}
		}

		// 加入新节点
		freq := 1
		node := &dLinkedNode{
			key:   key,
			value: value,
			freq:  freq,
			prev:  nil,
			next:  nil,
		}

		if oldHead := c.freqMap[freq]; oldHead != nil {
			oldHead.prev = node
			node.next = oldHead
		}

		c.valMap[key] = node
		c.freqMap[freq] = node

		if freq > c.maxFreq {
			c.maxFreq = freq
		}

		return
	}

	node := c.valMap[key]
	node.value = value
	// 已经是头节点
	if node.prev == nil {
		return
	}

	node.prev.next = node.next
	if node.next != nil {
		node.next.prev = node.prev
	}

	oldHead := c.freqMap[node.freq]
	oldHead.prev = node
	node.next = oldHead
	node.prev = nil
	c.freqMap[node.freq] = node
}
