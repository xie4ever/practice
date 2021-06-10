package sync

import (
	"sync"
	"testing"
)

type s struct {
	mutex sync.Mutex
}

func TestDeadlock(t *testing.T) {
	s1 := s{}
	s1.mutex.Lock()
	s2 := s1
	s1.mutex.Unlock()
	s2.mutex.Lock()
	s2.mutex.Unlock()
}

func TestDeadlock2(t *testing.T) {
	s1 := s{}
	s1.mutex.Lock()
	s1.mutex.Lock()
	s1.mutex.Unlock()
}
