package sorting

func MinSelectionSort(list []int) {
	for i := 0; i < len(list)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(list); j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}

func MaxSelectionSort(list []int) {
	for i := len(list) - 1; i > 0; i-- {
		maxIndex := i
		for j := i - 1; j >= 0; j-- {
			if list[maxIndex] < list[j] {
				maxIndex = j
			}
		}
		if maxIndex != i {
			list[i], list[maxIndex] = list[maxIndex], list[i]
		}
	}
}
