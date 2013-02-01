package edn

type Stack struct {
	top  *StackItem
	size int
}

type StackItem struct {
	value interface{}
	next  *StackItem
}

func (s *Stack) Push(value interface{}) {
	s.top = &StackItem{value, s.top}
	s.size++
}

func (s *Stack) Pop() (value interface{}) {
	if s.size == 0 {
		return
	}

	value = s.top.value
	s.top = s.top.next
	s.size--

	return value
}

func (s *Stack) Length() int {
	return s.size
}
