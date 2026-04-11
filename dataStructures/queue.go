package datastructures

import "errors"

type Queue[T any] struct {
	items []T
}

func(q *Queue[T]) EnQueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) DeQueue() (T, error){
	var elem T
	if len(q.items) == 0 {
		return elem, errors.New("Queue is empty")
	}
	elem = q.items[0]
	q.items = q.items[1:]
	return elem, nil
}

func (q *Queue[T]) Front()(T, error) {
	var elem T
	if len(q.items) == 0 {
		return elem, errors.New("Queue is empty")
	}
	return q.items[0], nil
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}


