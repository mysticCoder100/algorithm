package intro

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	search := 2
	count := 0

	expectedCount := 2
	expectedIndex := 1

	index := BinarySearch(list, search, &count)

	t.Logf("Found Index: %d, Total Operations: %d", index, count)

	if expectedCount != count {
		t.Errorf("Expecting the count to be %d found %d.", expectedCount, count)
	}

	if expectedIndex != index {
		t.Errorf("Expecting the index to be %d found %d.", expectedIndex, index)
	}
}

func TestFindSmallestIndex(t *testing.T) {
	list := []int{4, 6, 7, 8, 3, 4, 3, 1, 9}

	expectedIndex := 7

	index := FindSmallestIndex(list)

	t.Logf("Found index: %d for smallest", index)

	if expectedIndex != index {
		t.Errorf("Expected index to be %d found %d.", expectedIndex, index)
	}
}

func TestSelectionSort(t *testing.T) {
	list := []int{4, 6, 7, 9, 3, 4, 3, 1, 8}

	expectedFirst := 1
	expectedLast := 9

	if expectedFirst != list[0] {
		t.Errorf("Expected first to be %d found %d.", expectedFirst, list[0])
	}

	if expectedLast != list[len(list)-1] {
		t.Errorf("Expected last to be %d found %d.", expectedLast, list[len(list)-1])
	}
}
