package exercise

type AverageTest struct {
	name     string
	input    []int
	expected float64
}

type Sum2DTest struct {
	name     string
	input    [][]int
	expected int
}

func createSum2DTestCases() []Sum2DTest {
	return []Sum2DTest{
		{
			name: "standard 3x3 grid",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 45,
		},
		{
			name: "jagged grid (rows with varying lengths)",
			input: [][]int{
				{10, 20},
				{1, 2, 3, 4},
				{5},
			},
			expected: 45,
		},
		{
			name: "grid with negative numbers",
			input: [][]int{
				{-5, 10},
				{3, -2},
			},
			expected: 6,
		},
		{
			name: "single element grid",
			input: [][]int{
				{42},
			},
			expected: 42,
		},
		{
			name:     "completely empty 2D slice",
			input:    [][]int{},
			expected: 0,
		},
		{
			name: "2D slice with empty row slices",
			input: [][]int{
				{},
				{1, 2},
				{},
			},
			expected: 3,
		},
		{
			name: "nil row inside 2D slice",
			input: [][]int{
				nil,
				{5, 5},
			},
			expected: 10,
		},
	}
}

func createTestCases() []AverageTest {
	return []AverageTest{
		{
			name:     "exact integer average",
			input:    []int{1, 2, 3},
			expected: 2.0,
		},
		{
			name:     "fractional average",
			input:    []int{1, 2},
			expected: 1.5,
		},
		{
			name:     "empty slice edge case",
			input:    []int{},
			expected: 0.0,
		},
	}
}
