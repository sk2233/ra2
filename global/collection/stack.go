/*
@author: sk
@date: 2023/9/9
*/
package collection

type Stack[T any] struct {
	data  []T
	index int
}

func (s *Stack[T]) Pop() T {
	res := s.data[s.index]
	s.data = s.data[:s.index]
	s.index--
	return res
}

func (s *Stack[T]) Peek() T {
	return s.data[s.index]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.index < 0
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
	s.index++
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0), index: -1}
}
