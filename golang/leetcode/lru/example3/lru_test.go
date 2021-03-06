package cache

import (
	"fmt"
	"sync"
	"testing"
)

// Test ...
func Test(t *testing.T) {
	c, _ := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}

// Test2 ...
func Test2(t *testing.T) {
	c, _ := Constructor(2)
	c.Put(1, 0)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}

// Test3 ...
func Test3(t *testing.T) {
	c, _ := Constructor(2)
	c.Put("a", "a")
	c.Put("b", "b")
	fmt.Println(c.Get("a"))
	c.Put("c", "c")
	fmt.Println(c.Get("b"))
	c.Put("d", "d")
	fmt.Println(c.Get("a"))
	fmt.Println(c.Get("c"))
	fmt.Println(c.Get("d"))
}

func testPut(waitGroup sync.WaitGroup, c *LRUCache) {
	go func(waitGroup sync.WaitGroup, c *LRUCache) {
		waitGroup.Add(1)
		defer waitGroup.Done()
		for i := 0; i < 10; i++ {
			c.Put(i, i)
		}
	}(waitGroup, c)
}

// Test4 ...
func Test4(t *testing.T) {
	var waitGroup sync.WaitGroup
	c, _ := Constructor(10)
	for i := 0; i < 100; i++ {
		testPut(waitGroup, c)
	}
	waitGroup.Wait()
	for k, v := range c.cacheMap {
		fmt.Println(fmt.Sprintf("k=%v, v=%v", k, v))
	}
}
