package cache

import (
	"fmt"
	"testing"
)

// Test ...
func Test(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}

// Test2 ...
func Test2(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	c.Put(3, 3)
	fmt.Println(c.Get(3))
}

// Test3 ...
func Test3(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1))
	c.Put(3, 3)
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
	c.Put(4, 4)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(3))
	fmt.Println(c.Get(4))
}

// Test4 ...
func Test4(t *testing.T) {
	c := Constructor(0)
	c.Put(0, 0)
	fmt.Println(c.Get(0))
}

// Test5 ...
func Test5(t *testing.T) {
	c := Constructor(1)
	c.Put(2, 1)
	fmt.Println(c.Get(2))
	c.Put(3, 2)
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(3))
}

// Test6 ...
func Test6(t *testing.T) {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
}

// Test7 ...
func Test7(t *testing.T) {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(4, 1)
	fmt.Println(c.Get(1))
	fmt.Println(c.Get(2))
	fmt.Println(c.Get(4))
}
