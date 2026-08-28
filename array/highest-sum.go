package array

func HeighestSum(list []int) int {
	if len(list) <= 0 {
		return 0
	}

	highestSum := 0
	cummulativeSum := 0

	for _, v := range list {
		cummulativeSum += v

		if cummulativeSum <= 0 {
			cummulativeSum = 0
		}

		if cummulativeSum > highestSum {
			highestSum = cummulativeSum
		}
	}

	return highestSum
}
