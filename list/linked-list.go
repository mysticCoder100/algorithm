package list

import "fmt"

type Node struct {
	data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	size int
}

func (l *LinkedList) getSize() int {
	return l.size
}

func (l *LinkedList) isEmpty() bool {
	return (l.size == 0)
}

func (l *LinkedList) AddHead(d int) {
	l.Head = &Node{data: d, Next: l.Head}
	l.size++
}

func (l *LinkedList) AddTail(d int) {
	node := &Node{d, nil}
	curr := l.Head
	l.size++

	if curr == nil {
		l.Head = node
		return
	}

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = node
}

func (l *LinkedList) Print() {
	curr := l.Head

	for curr != nil {
		fmt.Print(curr.data, ", ")
		curr = curr.Next
	}
	fmt.Println()
}

func (l *LinkedList) SortedInsertion(d int) {
	newNode := &Node{data: d}
	curr := l.Head

	if curr == nil || curr.data > d {
		newNode.Next = curr
		l.Head = newNode
		return
	}

	for curr.Next != nil && curr.Next.data < d {
		curr = curr.Next
	}

	newNode.Next = curr.Next
	curr.Next = newNode

}
