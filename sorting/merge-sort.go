package sorting

func MergeSort(list []int, start int, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2

	MergeSort(list, start, mid)
	MergeSort(list, mid+1, end)
	merge(list, start, end, mid)
}

func merge(list []int, start int, end int, mid int) {
	s := start
	m := mid + 1
	l := make([]int, 0)

	for s <= mid && m <= end {
		if list[s] < list[m] {
			l = append(l, list[s])
			s++
		} else {
			l = append(l, list[m])
			m++
		}
	}

	for s <= mid {
		l = append(l, list[s])
		s++
	}

	for m <= end {
		l = append(l, list[m])
		m++
	}

	for i := start; i <= end; i++ {
		list[i] = l[i-start]
	}
}
