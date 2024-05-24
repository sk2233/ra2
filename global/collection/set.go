/*
@author: sk
@date: 2023/8/6
*/
package collection

type Set[T comparable] struct {
	data map[T]struct{}
}

func NewSet[T comparable](ts ...T) *Set[T] {
	res := &Set[T]{data: make(map[T]struct{})}
	for _, t := range ts {
		res.Add(t)
	}
	return res
}

func (s *Set[T]) Add(value T) {
	s.data[value] = struct{}{}
}

func (s *Set[T]) AddAll(values ...T) {
	for _, value := range values {
		s.Add(value)
	}
}

func (s *Set[T]) Has(value T) bool {
	_, ok := s.data[value]
	return ok
}
