package main

import (
	"fmt"

	"github.com/mysticCoder100/algorithm/list"
)

func main() {
	list := &list.LinkedList{}
	list.AddHead(2)
	list.AddHead(5)
	list.AddHead(6)
	fmt.Println(list.Head.Next.Next)
}
