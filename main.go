package main

import (
	"fmt"

	"github.com/mysticCoder100/algorithm/sorting"
)

func main() {
	list := []int{5, 4, 3, 2, 1, 9, 8, 7}
	fmt.Println(list)
	sorting.MergeSort(list, 0, len(list)-1)
	fmt.Println(list)
}
