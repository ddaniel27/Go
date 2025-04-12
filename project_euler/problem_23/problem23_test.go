package problem23

import "testing"

func TestProblem23_Func(t *testing.T) {
	tests := []struct {
		name   string
		output uint
	}{
		{"Problem 23 - Non-abundant sums", 4179871},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem23()
			if got != test.output {
				t.Errorf("Problem23() = got %v, want %v", got, test.output)
			}
		})
	}
}

func BenchmarkProblem23_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem23()
	}
}
