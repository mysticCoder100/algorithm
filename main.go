package main

import (
	"fmt"

	"github.com/mysticCoder100/algorithm/sorting"
)

func main() {
	list := []int{8, 1, 9, 2, 7, 3, 6, 4, 5}
	fmt.Println(list)
	item := sorting.QuickSelect(list, 0, len(list)-1, 5)
	fmt.Println(item)
}
