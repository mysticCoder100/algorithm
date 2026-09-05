package main

import (
	"fmt"

	"github.com/mysticCoder100/algorithm/sorting"
)

func main() {
	list := []int{5, 4, 3, 2, 1, 9, 8, 7, 6}
	fmt.Println(list)
	sorting.QuickSort(list, 0, len(list)-1)
	fmt.Println(list)
}
