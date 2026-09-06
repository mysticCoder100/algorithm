package sorting

func QuickSort(list []int, start int, end int) {
	if start >= end {
		return
	}

	partition := partition(list, start, end)

	QuickSort(list, start, partition-1)
	QuickSort(list, partition+1, end)

}

// the k passed is 1 based position
func QuickSelect(list []int, start int, end int, k int) int {
	if k < 0 || k > end-start+1 {
		return -1
	}

	partition := partition(list, start, end)

	if partition-start+1 == k {
		return list[partition]
	} else if partition-start+1 > k {
		return QuickSelect(list, start, partition-1, k)
	} else {
		return QuickSelect(list, partition+1, end, k-partition+start+1)
	}

}

func partition(list []int, start int, end int) int {
	pivot := list[end]

	i := start

	for j := start; j < end; j++ {
		if list[j] < pivot {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}

	list[i], list[end] = list[end], list[i]
	return i
}
