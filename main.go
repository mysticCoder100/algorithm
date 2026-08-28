package main

import (
	"fmt"

	"github.com/mysticCoder100/algorithm/array"
)

func main() {
	// list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := []int{1, -2, 3, 4, -4, 6, -14, 8, 2}
	fmt.Println(array.HeighestSum(list))
}
