package recursion

import (
	"testing"
)

func TestCountDown(t *testing.T) {
	time := 3
	Countdown(time)
}

func TestQuickSort(t *testing.T) {
	list := []int{2, 6, 8, 5, 4}

	expectedFirstItem := 2
	expectedLastItem := 8

	QuickSort(list, 0, len(list)-1)

	if list[0] != expectedFirstItem {
		t.Errorf("Expected first item to be %d, found %d", expectedFirstItem, list[0])
	}

	if list[len(list)-1] != expectedLastItem {
		t.Errorf("Expected last item to be %d, found %d", expectedLastItem, list[len(list)-1])
	}
}
