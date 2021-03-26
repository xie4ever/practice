package sort

import (
	"fmt"
	"testing"
)

// TestQuickSort ...
func TestQuickSort(t *testing.T) {
	var arr = []int{5, 4, 8, 3, 2, 9, 1, 7, 6}
	quickSort(arr, 0, len(arr)-1)
	for _, a := range arr {
		fmt.Println(a)
	}
}

func quickSort(arr []int, low, high int) {
	if len(arr) <= 0 {
		return
	}
	if low >= high {
		return
	}

	var (
		i     = low
		j     = high
		index = arr[i]
	)

	for i < j {
		for i < j && arr[j] >= index {
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}
		for i < j && arr[i] < index {
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = index
	quickSort(arr, low, i-1)
	quickSort(arr, i+1, high)
}
