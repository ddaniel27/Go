package problem22

import "testing"

func TestProblem22_Func(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output int
	}{
		{"Problem 22 - Names scores", "\"COLIN\"", 53},            // 3 + 15 + 12 + 9 + 14 = 53
		{"Problem 22 - Names scores", "\"COLIN\",\"COLIN\"", 159}, // 53*1 + 53*2 = 159
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem22(test.input)
			if got != test.output {
				t.Errorf("Problem22() = got %v, want %v", got, test.output)
			}
		})
	}
}

func BenchmarkProblem22_Func(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Problem22("\"COLIN\",\"COLIN\"")
	}
}
