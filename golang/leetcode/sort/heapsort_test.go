package sort

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func lastStoneWeight(stones []int) int {
	h := &hp{stones}
	heap.Init(h)
	for h.Len() > 1 {
		a := h.pop()
		b := h.pop()
		if a > b {
			c := a - b
			h.push(c)
		}
	}
	if h.Len() > 0 {
		return h.IntSlice[0]
	}
	return 0
}

// TestHeapSort ...
func TestHeapSort(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(lastStoneWeight(arr))
}
