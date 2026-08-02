package recursion

import "fmt"

func Countdown(num int) {
	fmt.Println(num)
	if num <= 1 {
		return
	}
	num--
	Countdown(num)
}

func RecursiveSum(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return list[0] + RecursiveSum(list[1:])
}

func RecursiveCount(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return 1 + RecursiveCount(list[1:])
}

func RecursiveFindMax(list []int) int {
	if len(list) == 0 {
		return 0
	}

	if len(list) == 1 {
		return list[0]
	}

	maxOfRest := RecursiveFindMax(list[1:])

	if list[0] > maxOfRest {
		return list[0]
	}

	return maxOfRest
}

func RecursiveBinarySearch(list []int, low int, high int, search int) int {

	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if list[mid] == search {
		return mid
	}

	if list[mid] > search {
		return RecursiveBinarySearch(list, low, mid-1, search)
	}

	return RecursiveBinarySearch(list, mid+1, high, search)
}

func QuickSort(list []int, low, high int) {
	if low < high {
		pivotIndex := partition(list, low, high)
		QuickSort(list, low, pivotIndex-1)
		QuickSort(list, pivotIndex+1, high)
	}
}

func partition(list []int, low, high int) int {
	pivot := list[high]

	i := low - 1
	var j int

	for j = low; j < high; j++ {
		if pivot >= list[j] {
			i++
			list[j], list[i] = list[i], list[j]
		}
	}

	list[i+1], list[j] = list[j], list[i+1]

	return i + 1
}
