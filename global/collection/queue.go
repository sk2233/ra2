/*
@author: sk
@date: 2023/9/9
*/
package collection

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Offer(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Poll() T {
	res := q.data[0]
	q.data = q.data[1:]
	return res
}

func (q *Queue[T]) Peek() T {
	return q.data[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Count() int {
	return len(q.data)
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{data: make([]T, 0)}
}
