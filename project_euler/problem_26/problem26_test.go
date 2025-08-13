package problem26

import "testing"

func TestA051626_Func(t *testing.T) {
	tests := []struct {
		name   string
		input  uint64
		output uint64
	}{
		{"Problem 26 - Length of recurring cycle for 7", 7, 6},
		{"Problem 26 - Length of recurring cycle for 3", 3, 1},
		{"Problem 26 - Length of recurring cycle for 11", 11, 2},
		{"Problem 26 - Length of recurring cycle for 13", 13, 6},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := A051626(test.input)
			if got != test.output {
				t.Errorf("A051626() = got %v, want %v", got, test.output)
			}
		})
	}
}

func TestProblem26_Func(t *testing.T) {
	tests := []struct {
		name   string
		input  uint64
		output uint64
	}{
		{"Problem 26 - d < 10", 10, 7},
		{"Problem 26 - d < 100", 100, 97},
		{"Problem 26 - d < 1000", 1000, 983},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Problem26(test.input)
			if got != test.output {
				t.Errorf("Problem26() = got %v, want %v", got, test.output)
			}
		})
	}
}
