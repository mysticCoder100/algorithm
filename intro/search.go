package intro

func BinarySearch(list []int, search int, count *int) int {
	mid, high, low := 0, len(list)-1, 0

	for low <= high {
		*count += 1
		mid = low + (high-low)/2

		if list[mid] == search {
			return mid
		} else if list[mid] < search {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func FindSmallestIndex(list []int) int {
	smallestIndex := 0

	for i, v := range list[1:] {
		if v < list[smallestIndex] {
			smallestIndex = i + 1
		}
	}

	return smallestIndex
}

func SelectionSort(list []int) {
	for i := 0; i < len(list); i++ {
		smallestIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[smallestIndex] {
				smallestIndex = j
			}
		}

		list[smallestIndex], list[i] = list[i], list[smallestIndex]
	}
}
