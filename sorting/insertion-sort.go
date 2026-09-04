package sorting

func InsertionSort(list []int) {
	size := len(list) - 1

	var j, temp int
	for i := 1; i <= size; i++ {
		temp = list[i]
		for j = i; j > 0 && list[j-1] > temp; j-- {
			list[j] = list[j-1]
		}
		list[j] = temp
	}
}

func RearInsertionSort(list []int) {
	size := len(list) - 1

	var j, temp int
	for i := size - 1; i >= 0; i-- {
		temp = list[i]
		for j = i; j < size && list[j+1] < temp; j++ {
			list[j] = list[j+1]
		}
		list[j] = temp
	}
}
