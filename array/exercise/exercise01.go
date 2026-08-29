package exercise

import "math"

/**
 * Find average of all the elements in a list.
 */
func AverageOfAllElements(list []int) float64 {
	if len(list) <= 0 {
		return 0.0
	}

	sum := 0

	for _, v := range list {
		sum += v
	}

	return float64(sum) / float64(len(list))
}

/**
 * Find the sum of all the elements of a two dimensional list.
 */
func SumOf2DList(list [][]int) int {
	if len(list) <= 0 {
		return 0
	}

	sum := 0

	for _, v := range list {
		for _, i := range v {
			sum += i
		}
	}

	return sum
}

/**
 * Find the second largest number in the list.
 */
func FindSecondLargest(list []int) int {
	if len(list) < 2 {
		return -1
	}

	largest := math.MinInt
	secondLargest := math.MinInt

	for _, current := range list {
		if current > largest {
			secondLargest = largest
			largest = current
		} else if current > secondLargest && current != largest {
			secondLargest = current
		}
	}

	if secondLargest == math.MinInt {
		return -1
	}

	return secondLargest
}

func SortZeroOne(list []int) {
	start := 0
	end := len(list) - 1

	for start < end {

		if list[start] == 0 {
			start += 1
		} else if list[end] == 1 {
			end -= 1
		} else {
			list[start], list[end] = list[end], list[start]
			end -= 1
			start += 1
		}
	}
}

func SortZeroOneTwo(list []int) {
	low := 0
	mid := 0
	high := len(list) - 1

	// list := []int{2, 1, 2, 2, 1, 0, 1, 0, 2, 0, 2, 2, 1}
	for mid <= high {
		switch list[mid] {
		case 0:
			list[low], list[mid] = list[mid], list[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			list[mid], list[high] = list[high], list[mid]
			high--
		}
	}

}

/**
 * Will implement for 4 elements.
 */
