package numeric

import "testing"

func TestLcmBase(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{
			3,
			6,
			6,
		},
		{
			1,
			12,
			12,
		},
		{
			12,
			12,
			12,
		},
		{
			352342100,
			34343659132,
			3025179245063264300,
		},
		// {
		// 	0,
		// 	0,
		// 	0,
		// },
		// {
		// 	12,
		// 	12,
		// 	1,
		// },
	}

	for _, tt := range tests {
		evaluated := lcmBase(tt.a, tt.b)
		if evaluated != tt.expected {
			t.Errorf("[Error] lcm(%d,%d) is Error. got=%d", tt.a, tt.b, evaluated)
		}
	}
}

func TestLcm(t *testing.T) {
	tests := []struct {
		inputs   []int
		expected int
	}{
		{
			[]int{3, 6, 9},
			18,
		},
		{
			[]int{120, 50, 5},
			600,
		},
		{
			[]int{64, 58, 9},
			16704,
		},
		{
			[]int{1, 1, 1},
			1,
		},
	}

	for _, tt := range tests {
		evaluated := Lcm(tt.inputs[0], tt.inputs[1], tt.inputs[2])
		if evaluated != tt.expected {
			t.Errorf("[Error] Lcm(%d,%d, %d) is Error. got=%d", tt.inputs[0], tt.inputs[1], tt.inputs[2], evaluated)
		}
	}
}
