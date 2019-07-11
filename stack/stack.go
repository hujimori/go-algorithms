package stack

type StackData struct {
	data interface{}
	next *StackData
}

type Stack struct {
	sp  *StackData
	top int
}

func NewStack() *Stack {
	return &Stack{top: 0}
}

func (s *Stack) Push(data interface{}) {
	s.sp = &StackData{data: data, next: s.sp}
	s.top++
}

func (s *Stack) Pop() interface{} {
	if s.top > 0 {
		item := s.sp.data
		s.sp = s.sp.next
		s.top--
		return item
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.top > 0 {
		return s.sp.data
	}
	return nil
}

func (s *Stack) IsEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return s.top
}
