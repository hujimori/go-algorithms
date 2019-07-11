package numeric

import "testing"

func TestGcdBase(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{
			3,
			6,
			3,
		},
		{
			1,
			12,
			1,
		},
		{
			12,
			12,
			12,
		},
		{
			352342100,
			34343659132,
			4,
		},
		{
			0,
			0,
			0,
		},
		// {
		// 	12,
		// 	12,
		// 	1,
		// },
	}

	for _, tt := range tests {
		evaluated := gcdBase(tt.a, tt.b)
		if evaluated != tt.expected {
			t.Errorf("[Error] gcd(%d,%d) is Error. got=%d", tt.a, tt.b, evaluated)
		}
	}
}

func TestGcd(t *testing.T) {
	tests := []struct {
		inputs   []int
		expected int
	}{
		{
			[]int{3, 6, 9},
			3,
		},
		{
			[]int{120, 50, 5},
			5,
		},
		{
			[]int{64, 58, 9},
			1,
		},
	}

	for _, tt := range tests {
		evaluated := Gcd(tt.inputs[0], tt.inputs[1], tt.inputs[2])
		if evaluated != tt.expected {
			t.Errorf("[Error] gcd(%d,%d, %d) is Error. got=%d", tt.inputs[0], tt.inputs[1], tt.inputs[2], evaluated)
		}
	}
}
