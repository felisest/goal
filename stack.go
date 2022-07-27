package goal

// Stack represents a instance of the stack data container.
type Stack[T any] struct {
	buff []T
	len  int
	cap  int
}

// MakeStack constructs and return a new Stack.
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

// Push inserts element at the top.
func (s *Stack[T]) Push(v T) {

	if s.len >= len(s.buff) {
		s.resize()
	}

	s.buff[s.len] = v
	s.len++
}

// Pop return the top element.
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

// Erase erases all unused elements.
func (s *Stack[T]) Erase() {
	new_buff := make([]T, s.cap)
	copy(new_buff, s.buff[:s.len])
	s.buff = new_buff
}

// Clear clears the container. After this call, Len() returns zero.
func (s *Stack[T]) Clear() {
	new_buff := make([]T, s.cap)
	s.buff = new_buff
}

// Len returns the number of elements.
func (s *Stack[T]) Len() int {
	return s.len
}

// Cap returns the capacity of container.
func (s *Stack[T]) Cap() int {
	return s.cap
}
