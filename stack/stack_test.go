package stack

type stack struct {
	arr []int
}

func (s *stack) len() int {
	return len(s.arr)
}

func (s *stack) push(val int) {
	s.arr = append(s.arr, val)
}

func (s *stack) top() int {
	return s.arr[len(s.arr)-1]
}

func (s *stack) bottom() int {
	return s.arr[0]
}

func (s *stack) pop() int {
	val := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return val
}
