package dynamic

import (
	"fmt"
	"testing"
)

// TestUpStair ...
func TestUpStairs(t *testing.T) {
	fmt.Println(upstairs(10))
}

func upstairs(step int) int {
	if step == 1 {
		return 1
	}
	if step == 2 {
		return 2
	}
	return upstairs(step-1) + upstairs(step-2)
}

func upstairs2(step int) int {
	if step == 1 {
		return 1
	}
	if step == 2 {
		return 2
	}
	var (
		a   = 1
		b   = 2
		tmp = 0
	)
	for i := 3; i <= step; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}

// TestUpStairs2 ...
func TestUpStairs2(t *testing.T) {
	fmt.Println(upstairs2(10))
}

func upstairs3(step int) int {
	if step == 1 {
		return 1
	}
	if step == 2 {
		return 2
	}
	var (
		a   = 1
		b   = 2
		tmp = 0
	)
	for i := 3; i < step; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}

// TestUpStairs3 ...
func TestUpStairs3(t *testing.T) {
	fmt.Println(upstairs3(10))
}
