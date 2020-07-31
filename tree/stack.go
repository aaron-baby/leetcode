package tree

import (
	"errors"
	"sync"
)

type IntStack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s    []int
}

func NewStack() *IntStack {
	return &IntStack{sync.Mutex{}, make([]int, 0)}
}

func (s *IntStack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.s = append(s.s, v)
}

func (s *IntStack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	l := len(s.s)
	if l == 0 {
		return 0, errors.New("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}
