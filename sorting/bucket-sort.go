package sorting

import "fmt"

func BucketSort() {
	// size 10
	list := []float32{0.78, 0.17, 0.39, 0.26, 0.72, 0.94, 0.21, 0.12, 0.23, 0.68}

	fmt.Println(list)
	bucketList := make([][]float32, len(list))

	for i := 0; i < len(list); i++ {
		key := int(list[i] * float32(len(list)))
		bucketList[key] = append(bucketList[key], list[i])
	}

	for _, v := range bucketList {
		insertionSort(v)
	}

	index := 0
	for _, v := range bucketList {
		for _, v := range v {
			list[index] = v
			index++
		}
	}

	fmt.Println(list)
}

func insertionSort(list []float32) {
	for i := 0; i < len(list); i++ {
		curr := list[i]
		var j int
		for j = i; j > 0 && list[j-1] > curr; j-- {
			list[j] = list[j-1]
		}

		list[j] = curr
	}
}
