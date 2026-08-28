package array

func Rotate(list []int, k int) {
	n := len(list)
	if n <= 1 {
		return
	}

	k = k % n
	if k < 0 {
		k = n + k
	}

	reverse(list, 0, k-1)
	reverse(list, k, n-1)
	reverse(list, 0, n-1)
}

func reverse(list []int, start int, end int) {
	for start < end {
		list[start], list[end] = list[end], list[start]
		start += 1
		end -= 1
	}
}
