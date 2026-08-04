package graph

import "errors"

type MyQueue struct {
	items []string
}

func (q *MyQueue) Enqueue(item string) {
	q.items = append(q.items, item)
}

func (q *MyQueue) Dequeue() (string, error) {
	if q.isEmpty() {
		return "", errors.New("The queue is empty")
	}

	dequeued := q.items[0]
	q.items[0] = ""
	q.items = q.items[1:]
	return dequeued, nil
}

func (q *MyQueue) isEmpty() bool {
	return len(q.items) == 0
}
