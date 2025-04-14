package problem24

import "testing"

func TestProblem24_Func(t *testing.T) {
	tests := []struct {
		name   string
		input  uint
		output string
	}{
		{"Problem 24 - Item in position 1", 1, "0123456789"},
		{"Problem 24 - Item in position 999999", 999999, "2783915406"},
		{"Problem 24 - Item in position 1000000", 1000000, "2783915460"},
		{"Problem 24 - Item in position 1000001", 1000001, "2783915604"},
		{"Problem 24 - Item in position 3628800", 3628800, "9876543210"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem24(test.input)
			if got != test.output {
				t.Errorf("Problem24() = got %v, want %v", got, test.output)
			}
		})
	}
}

func BenchmarkProblem24_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem24(1000000)
	}
}
