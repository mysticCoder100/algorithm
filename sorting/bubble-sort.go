package sorting

func BubbleSort(list []int) {
	size := len(list) - 1

	for i := 0; i <= size; i++ { // pointer
		for j := 0; j < size-i; j++ { // iterator
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func ImprovedBubbleSort(list []int) {
	size := len(list) - 1
	swapped := true
	for i := 0; i <= size && swapped; i++ { // pointer
		swapped = false
		for j := 0; j < size-i; j++ { // iterator
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
	}
}
