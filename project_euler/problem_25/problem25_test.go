package problem25

import "testing"

func TestProblem25_Func(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"Problem 25 - Index of first term with 3 digits", 3, 12},
		{"Problem 25 - Index of first term with 1000 digits", 1000, 4782},
		{"Problem 25 - Index of first term with 100 digits", 100, 476},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem25(test.input)
			if got != test.output {
				t.Errorf("Problem25() = got %v, want %v", got, test.output)
			}
		})
	}
}

func BenchmarkProblem25_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem25(1000)
	}
}
