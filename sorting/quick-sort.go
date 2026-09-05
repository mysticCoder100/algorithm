package sorting

func QuickSort(list []int, start int, end int) {
	if start >= end {
		return
	}

	partition := partition(list, start, end)

	QuickSort(list, start, partition-1)
	QuickSort(list, partition+1, end)

}

func partition(list []int, start int, end int) int {
	pivot := list[end]

	i := start - 1

	for j := start; j < end; j++ {
		if list[j] < pivot {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	i++
	list[i], list[end] = list[end], list[i]
	return i
}
