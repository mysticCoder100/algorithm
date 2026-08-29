package exercise

import "testing"

func TestAverageOfAllElements(t *testing.T) {

	testCases := createTestCases()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := AverageOfAllElements(tt.input)
			if got != tt.expected {
				t.Errorf("Expected %f, Got %f", tt.expected, got)
			}
		})
	}
}

func TestSum2D(t *testing.T) {
	testCases := createSum2DTestCases()

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			sum := tt.expected
			if sum != tt.expected {
				t.Errorf("Expected %d, Got %d", tt.expected, sum)
			}
		})
	}
}
