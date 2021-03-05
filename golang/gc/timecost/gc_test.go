package timecost

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// from https://www.wolai.com/mayj/rLGcCeJpM6Jb2JvSfpTru4

const (
	objCount  = 1 * 10000 // 对象个数
	testCount = 10        // 次数
)

type obj struct {
	data [1024 * 1024]byte // 1MB
}

func init() {
	fmt.Println(runtime.Version())
}

func Test_gc_slice(t *testing.T) {
	var (
		sum time.Duration
		max time.Duration
		min time.Duration
	)
	for i := 0; i < testCount; i++ {
		d := baseTestSlice()
		sum += d
		if i == 0 {
			max = d
			min = d
		} else {
			if max < d {
				max = d
			}
			if min > d {
				min = d
			}
		}
	}
	fmt.Println("gc_slice -c", testCount, "avg", sum/testCount, "max", max, "min", min)
}

func baseTestSlice() time.Duration {
	l := make([]*obj, 0)
	for i := 0; i < objCount; i++ {
		l = append(l, &obj{})
	}
	ti := time.Now()
	l = nil
	runtime.GC()
	return time.Since(ti)
}

func Test_gc_map(t *testing.T) {
	var (
		sum time.Duration
		max time.Duration
		min time.Duration
	)
	for i := 0; i < testCount; i++ {
		d := baseTestMap()
		sum += d
		if i == 0 {
			max = d
			min = d
		} else {
			if max < d {
				max = d
			}
			if min > d {
				min = d
			}
		}
	}
	fmt.Println("gc_map -c", testCount, "avg", sum/testCount, "max", max, "min", min)
}

func baseTestMap() time.Duration {
	m := make(map[int]*obj)
	for i := 0; i < objCount; i++ {
		m[i] = &obj{}
	}
	ti := time.Now()
	m = nil
	runtime.GC()
	return time.Since(ti)
}
