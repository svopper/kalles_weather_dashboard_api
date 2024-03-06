package metObs

import (
	"testing"
)

func TestRemoveMinAndMaxValue(t *testing.T) {
	var tests = []struct {
		input    []float64
		expected []float64
	}{
		{
			[]float64{
				1.0,
				2.0,
				3.0,
				4.0,
			},
			[]float64{
				2.0,
				3.0,
			},
		},
		{
			[]float64{
				1.0,
				1.0,
				2.0,
				3.0,
				3.0,
			},
			[]float64{
				1.0,
				2.0,
				3.0,
			},
		},
		{
			[]float64{
				1.0,
				1.0,
				1.0,
				1.0,
			},
			[]float64{
				1.0,
				1.0,
			},
		},
	}

	for _, test := range tests {
		if output := removeMinAndMaxValue(test.input); !equal(output, test.expected) {
			t.Errorf("removeMinAndMaxValue(%v) = %v, want %v", test.input, output, test.expected)
		}
	}
}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
