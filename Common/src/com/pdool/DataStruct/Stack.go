package DataStruct

import (
	"fmt"
)

type Stack struct {
	items []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.items = append(s.items, element)
}

func (s *Stack) Pop() interface{} {
	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem
}

func (s *Stack) Peek() interface{} {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = []interface{}{}
}

func (s *Stack) Print() {
	fmt.Printf("%#v\n", s)
}
