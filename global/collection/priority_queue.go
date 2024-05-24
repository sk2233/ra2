/*
@author: sk
@date: 2023/9/30
*/
package collection

type TComparator[T any] func(T, T) bool

type PriorityQueue[T any] struct {
	items      []T
	comparator TComparator[T]
}

func NewPriorityQueue[T any](comparator TComparator[T]) *PriorityQueue[T] {
	return &PriorityQueue[T]{comparator: comparator, items: make([]T, 0)}
}

func (p *PriorityQueue[T]) Offer(item T) {
	p.items = append(p.items, item)
	p.swim(len(p.items) - 1)
}

func (p *PriorityQueue[T]) Poll() T {
	res := p.items[0]
	newLen := len(p.items) - 1
	p.items[0] = p.items[newLen]
	p.items = p.items[:newLen]
	p.sink(0)
	return res
}

func (p *PriorityQueue[T]) Peek() T {
	return p.items[0]
}

func (p *PriorityQueue[T]) IsEmpty() bool {
	return len(p.items) == 0
}

func (p *PriorityQueue[T]) swim(index int) { // 上浮
	if index == 0 {
		return
	}
	parent := (index - 1) / 2
	if p.comparator(p.items[index], p.items[parent]) {
		p.items[index], p.items[parent] = p.items[parent], p.items[index]
		p.swim(parent)
	}
}

func (p *PriorityQueue[T]) sink(index int) {
	l := len(p.items)
	tarIndex := index
	if index*2+1 < l && p.comparator(p.items[index*2+1], p.items[tarIndex]) {
		tarIndex = index*2 + 1
	}
	if index*2+2 < l && p.comparator(p.items[index*2+2], p.items[tarIndex]) {
		tarIndex = index*2 + 2
	}
	if tarIndex != index {
		p.items[index], p.items[tarIndex] = p.items[tarIndex], p.items[index]
		p.sink(tarIndex)
	}
}
