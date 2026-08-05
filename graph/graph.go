package graph

import (
	"fmt"
)

var MyGraph map[string][]string = map[string][]string{
	"you":    {"alice", "bob", "claire"},
	"bob":    {"anuj", "peggy"},
	"alice":  {"peggy"},
	"claire": {"thom", "jonny"},
	"anuj":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

func Perform(item []string) bool {
	queue := MyQueue[string]{items: item}
	searched := make(map[string]bool)
	searched["you"] = true

	for !queue.IsEmpty() {
		item, err := queue.Dequeue()

		if err != nil {
			fmt.Println(err.Error())
			break
		}

		if searched[item] {
			continue
		}

		searched[item] = true

		if isPersonSeller(item) {
			fmt.Printf("%s, is a mango seller", item)
			return true
		}

		for _, v := range MyGraph[item] {
			if !searched[v] {
				searched[v] = true
				queue.Enqueue(v)
			}
		}

	}

	return false
}

func isPersonSeller(name string) bool {
	return name[len(name)-1] == 'm'
}
