package queue

import "testing"

var q *Queue
var tests = []int{
	1,
	2,
	3,
	4,
	5,
}

func testInit() {
	q = NewQueue()
	for _, tt := range tests {
		q.Enqueue(tt)
	}
}

func TestDequeue(t *testing.T) {
	testInit()

	for _, tt := range tests {
		var d interface{}
		if d = q.Dequeue(); d != tt {
			t.Errorf("Dequeue() is error. got=%d. expected=%d", d, tt)
		}
		t.Log(d)
	}
}

func TestFront(t *testing.T) {
	testInit()

	if q.Front() != tests[0] {
		t.Errorf("Front() is error. got=%d. expected=%d", q.Front(), tests[0])
	}

}

func TestLast(t *testing.T) {
	testInit()

	if q.Last() != tests[len(tests)-1] {
		t.Errorf("Last() is error. got=%d. exptected=%d", q.Last(), tests[len(tests)-1])
	}
}

func TestIsEmpty(t *testing.T) {
	testInit()

	if q.IsEmpty() {
		t.Errorf("IsEmpty() is error. got=%t.", q.IsEmpty())
	}
}

func TestSize(t *testing.T) {
	testInit()

	if q.Size() != len(tests) {
		t.Errorf("Size() is error. got=%d. expected=%d.", q.Size(), len(tests))
	}
}
