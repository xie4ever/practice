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

// TestInsertSort2 ...
func TestInsertSort2(t *testing.T) {
	var arr = []int{12, 1, 3, 46, 5, 0, -3, 12, 35, 16}
	insertSort2(arr)
	for _, a := range arr {
		fmt.Println(a)
	}
}

func insertSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		if arr[i] >= arr[i-1] {
			continue
		} else {
			tmp := arr[i]
			j := i
			for ; j > 0; j-- {
				if arr[j-1] > tmp {
					arr[j] = arr[j-1]
				} else {
					break
				}
			}
			arr[j] = tmp
		}
	}
}
