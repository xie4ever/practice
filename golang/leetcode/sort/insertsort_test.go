package sort

import (
	"fmt"
	"testing"
)

// TestInsertSort ...
func TestInsertSort(t *testing.T) {
	var arr = []int{12, 1, 3, 46, 5, 0, -3, 12, 35, 16}
	insertSort(arr)
	for _, a := range arr {
		fmt.Println(a)
	}
}

func insertSort(arr []int) {
	index := 0
	for index < len(arr)-1 {
		if arr[index] <= arr[index+1] {
			index++
			continue
		} else {
			i := index + 1
			tmp := arr[i]
			for ; i > 0; i-- {
				if arr[i-1] > tmp {
					arr[i] = arr[i-1]
				} else {
					break
				}
			}
			arr[i] = tmp
			index++
		}
	}
}
