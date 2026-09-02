package prac

import (
	"strconv"
)

func SortHusbandWifeList(list []string) {
	for i := 0; i < len(list)-1; i++ {
		lowestIndex := i
		for j := i + 1; j < len(list); j++ {
			if i == getPosition(list[j]) {
				lowestIndex = j
			}
		}
		list[i], list[lowestIndex] = list[lowestIndex], list[i]
	}
}

func getPosition(s string) int {
	userType := s[0]
	position, _ := strconv.Atoi(s[1:])

	if userType == 'H' {
		return 2 * (position - 1)
	} else {
		return 2*(position-1) + 1
	}
}
