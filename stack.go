package goal

type Stack[T any] struct {
	buff []T
	len  int
	cap  int
}

func MakeStack[T any](capacity int) *Stack[T] {

	return &Stack[T]{buff: make([]T, capacity), cap: capacity}
}

func (s *Stack[T]) resize() {

	new_capacity := grow_container(s.cap)
	new_buff := make([]T, new_capacity)

	copy(new_buff, s.buff[:s.len])

	s.buff = new_buff
	s.cap = new_capacity
}

func (s *Stack[T]) Push(v T) {

	if s.len >= len(s.buff) {
		s.resize()
	}

	s.buff[s.len] = v
	s.len++
}

func (s *Stack[T]) Pop() (T, bool) {

	var dr T
	if s.len == 0 {
		return dr, false
	}

	v := s.buff[s.len-1]
	s.buff[s.len-1] = dr
	s.len--

	return v, true
}

func (s *Stack[T]) Len() int {
	return s.len
}

func (s *Stack[T]) Cap() int {
	return s.cap
}
