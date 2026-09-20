package list

import "fmt"

type Node struct {
	data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

func (l *LinkedList) getSize() int {
	return l.Size
}

func (l *LinkedList) isEmpty() bool {
	return (l.Size == 0)
}

func (l *LinkedList) AddHead(d int) {
	l.Head = &Node{data: d, Next: l.Head}
	l.Size++
}

func (l *LinkedList) AddTail(d int) {
	node := &Node{d, nil}
	curr := l.Head
	l.Size++

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

func (l *LinkedList) Exists(d int) bool {
	curr := l.Head

	for curr != nil {
		if curr.data == d {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *LinkedList) Delete() (int, bool) {
	if l.isEmpty() {
		return 0, false
	}
	temp := l.Head
	l.Head = temp.Next
	l.Size--
	return temp.data, true
}

func (l *LinkedList) DeleteItem(d int) (int, bool) {
	if l.isEmpty() {
		return 0, false
	}

	if d == l.Head.data {
		l.Head = l.Head.Next
		l.Size--
		return d, true
	}

	curr := l.Head

	for curr.Next != nil {
		if curr.Next.data == d {
			curr.Next = curr.Next.Next
			l.Size--
			return d, true
		}
		curr = curr.Next
	}

	return 0, false
}

func (l *LinkedList) DeleteAllOccurence(d int) {
	curr := l.Head

	for curr != nil && curr.data == d {
		l.Head = curr.Next
		curr = curr.Next
		l.Size--
	}

	for curr != nil {
		if curr.Next != nil && curr.Next.data == d {
			curr.Next = curr.Next.Next
			l.Size--
		} else {
			curr = curr.Next
		}
	}
}

func (l *LinkedList) Free() {
	l.Head = nil
	l.Size = 0
}
