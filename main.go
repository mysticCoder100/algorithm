package main

import (
	"fmt"

	"github.com/mysticCoder100/algorithm/list"
)

func main() {
	list := &list.LinkedList{}
	list.SortedInsertion(2)
	list.SortedInsertion(6)
	list.SortedInsertion(5)
	list.SortedInsertion(1)
	fmt.Println(list.Head.Next)
}
