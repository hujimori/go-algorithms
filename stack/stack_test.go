package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	tests := []int{
		1,
		2,
		3,
		4,
		5,
	}

	s := NewStack()

	for _, tt := range tests {
		s.Push(tt)

		if tt != s.Peek() {
			t.Errorf("push(%d) is Error. got=%d", tt, s.Peek())
		}

		t.Log(s.Peek())

	}

	if int(s.Size()) != len(tests) {
		t.Errorf("Size() is Error. got=%d. expected=%d", s.Size(), len(tests))
	}

	for i := len(tests) - 1; i >= 0; i-- {
		x := s.Pop()
		t.Log(x)

		if x != tests[i] {
			t.Errorf("pop() is Error. got=%d. expected=%d", x, tests[i])
		}
	}

}
