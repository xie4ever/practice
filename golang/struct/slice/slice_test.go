package test

import (
	"fmt"
	"testing"
)

// Test ...
func Test(t *testing.T) {
	var s []string
	var str = "test"
	fmt.Println(str)
	fmt.Println(&str) // 0xc0000444e0
	s = append(s, str)
	fmt.Println(s[0])
	fmt.Println(&(s[0])) // 0xc000044500
	str = "fuck"
	fmt.Println(s[0])    // test
	fmt.Println(&(s[0])) // 0xc000044500
}

// Test2 ...
func Test2(t *testing.T) {
	var s []*string
	var str = "test"
	var ptr = &str
	fmt.Println(ptr)  // 0xc0000444f0
	fmt.Println(&ptr) // 0xc000006038
	s = append(s, ptr)
	fmt.Println(s[0])    // 0xc0000444f0
	fmt.Println(&(s[0])) // 0xc000006040
	str = "fuck"
	fmt.Println(s[0])    // 0xc0000444f0
	fmt.Println(&(s[0])) // 0xc000006040
}

// Test3 ...
func Test3(t *testing.T) {
	var str = "test"
	s := []string{str}
	fmt.Println(str)
	fmt.Println(&str) // 0xc000044500
	fmt.Println(s[0])
	fmt.Println(&(s[0])) // 0xc000044510
	str = "fuck"
	fmt.Println(s[0])    // test
	fmt.Println(&(s[0])) // 0xc000044510
}

// Test4 ...
func Test4(t *testing.T) {
	var str = "test"
	s := make([]string, 1, 1)
	s[0] = str
	fmt.Println(str)
	fmt.Println(&str) // 0xc000044500
	fmt.Println(s[0])
	fmt.Println(&(s[0])) // 0xc000044510
	str = "fuck"
	fmt.Println(s[0])    // test
	fmt.Println(&(s[0])) // 0xc000044510
}

// Test5 ...
func Test5(t *testing.T) {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := array[1:3]
	s2 := array[5:9]
	fmt.Println(s1)      // [1 2]
	fmt.Println(len(s1)) // 2
	fmt.Println(cap(s1)) // 9
	fmt.Println(s2)      // [5 6 7 8]
	fmt.Println(len(s2)) // 4
	fmt.Println(cap(s2)) // 5

	s2 = append(s2, 9, 10)
	fmt.Println(len(s2)) // 6
	fmt.Println(cap(s2)) // 10

	array[5] = 1    // 尝试修改array，看看s2的底层数组会不会改变
	fmt.Println(s2) // [5 6 7 8 9 10]
}

// Test6 ...
func Test6(t *testing.T) {
	s := make([]int, 2, 2)
	fmt.Println(s) // [0 0]
	fmt.Println(&s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s1 := append(s, 3, 4)

	fmt.Println()
	fmt.Println(s) // [0 0]
	fmt.Println()

	fmt.Println(s1) // [0 0 3 4]
	fmt.Println(&s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}
