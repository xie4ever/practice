package sort

import (
	"errors"
	"fmt"
	"testing"
)

type bitMap struct {
	bits []byte
}

// 初始化一个BitMap
// 一个byte有8位，可代表8个数字，除8后加1为存放最大数所需的容量
func newBitMap(max int) *bitMap {
	bits := make([]byte, (max/8)+1)
	return &bitMap{bits: bits}
}

// 1.找到所属的int，num / 8
// 2.找到在这个int中1的位移，num % 8
// 3.设置1，b[byteIndex] = b[byteIndex] | (1 << bitIndex)
func (b *bitMap) add(num int) error {
	byteIndex := num / 8
	if byteIndex >= len(b.bits) {
		return errors.New("num > max")
	}
	bitIndex := num % 8
	b.bits[byteIndex] = b.bits[byteIndex] | (1 << bitIndex)
	return nil
}

// 1.找到所属的int，num / 8
// 2.找到在这个int中1的位移，num % 8
// 3.设置0，b[byteIndex] = b[byteIndex] & (^(1 << bitIndex))
func (b *bitMap) clear(num int) error {
	byteIndex := num / 8
	if byteIndex >= len(b.bits) {
		return errors.New("num > max")
	}
	bitIndex := num % 8
	b.bits[byteIndex] = b.bits[byteIndex] & (^(1 << bitIndex))
	return nil
}

// 1.找到所属的int，num / 8
// 2.找到在这个int中1的位移，num % 8
// 3. b[byteIndex] & (1 << bitIndex)
func (b *bitMap) contain(num int) bool {
	byteIndex := num / 8
	if byteIndex >= len(b.bits) {
		return false
	}
	bitIndex := num % 8
	return !(b.bits[byteIndex]&(1<<bitIndex) == 0)
}

func (b *bitMap) list() []int {
	var dataList []int
	size := len(b.bits) * 8
	for i := 0; i < size; i++ {
		if b.contain(i) {
			dataList = append(dataList, i)
		}
	}
	return dataList
}

// TestBitMap ...
func TestBitMap(t *testing.T) {
	b := newBitMap(100)
	b.add(3)
	fmt.Println(b.contain(3))
	b.clear(3)
	fmt.Println(b.contain(3))
}

// TestList ...
func TestList(t *testing.T) {
	b := newBitMap(100)
	b.add(3)
	b.add(5)
	b.add(12)
	b.add(50)
	b.add(98)
	b.add(99)
	fmt.Println(b.list())
}
