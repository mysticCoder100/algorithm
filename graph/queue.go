package graph

import "errors"

type MyQueue[T any] struct {
	items []T
}

func (q *MyQueue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *MyQueue[T]) Dequeue() (T, error) {
	var zero T

	if q.IsEmpty() {
		return zero, errors.New("The queue is empty")
	}

	dequeued := q.items[0]
	q.items[0] = zero
	q.items = q.items[1:]
	return dequeued, nil
}

func (q *MyQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}
