package problem21

import "testing"

func TestProblem21_Func(t *testing.T) {
	tests := []struct {
		name     string
		input    uint
		expected uint
	}{
		{"Problem 21 - Amicable numbers", 10000, 31626},
		{"Problem 21 - Amicable numbers", 1000, 504},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem21(test.input)
			if got != test.expected {
				t.Errorf("Problem21() = got %v, want %v", got, test.expected)
			}
		})
	}
}

func BenchmarkProblem21_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem21(10000)
	}
}
