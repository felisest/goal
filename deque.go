package goal

// Deque represents a instance of the queue data container.
type Deque[T any] struct {
	
	front *node[T]
	back *node[T]
	len   int
}

type node[T any] struct {
	elem *T
	next *node[T]
	prev *node[T]
}

// MakeDeque constructs and return a new Deque.
func MakeDeque[T any]() *Deque[T] {

	return &Deque[T]{front: nil, back: nil, len: 0}
}

// PushBack adds an element to the end
func (dq *Deque[T]) PushBack(v T) {

	if dq.back == nil && dq.front == nil && dq.len == 0 {
		dq.back = &node[T]{elem: &v, next: nil, prev: nil}
		dq.front = dq.back
		dq.len = 1

	} else {
		curr_node := dq.back
		dq.back = &node[T]{elem: &v, next: nil, prev: curr_node}
		curr_node.next = dq.back
		dq.len++
	}
}

// PopBack returns the last element
func (dq *Deque[T]) PopBack() (T, bool) {

	var dr T

	if dq.back == nil && dq.front == nil && dq.len == 0 {
		return dr, false

	} else {
		ret_node := dq.back
		dq.back = ret_node.prev
		dr = *ret_node.elem
		dq.len--
	}
	return dr, true
}

// PushFront inserts an element to the beginning
func (dq *Deque[T]) PushFront(v T) {

	if dq.back == nil && dq.front == nil && dq.len == 0 {
		dq.back = &node[T]{elem: &v, next: nil, prev: nil}
		dq.front = dq.back
		dq.len = 1

	} else {
		curr_node := dq.front
		dq.front = &node[T]{elem: &v, next: curr_node, prev: nil}
		curr_node.prev = dq.front
		dq.len++
	}
}

// PopFront returns the first element
func (dq *Deque[T]) PopFront() (T, bool) {

	var dr T

	if dq.back == nil && dq.front == nil && dq.len == 0 {
		return dr, false

	} else {
		ret_node := dq.front
		dq.front = ret_node.next
		dr = *ret_node.elem
		dq.len--
	}
	return dr, true
}

// Clear clears the container. After this call, Len() returns zero.
func (dq *Deque[T]) Clear() {
	dq.len = 0
	dq.front = nil
	dq.back = nil
}

// Len returns the number of elements.
func (dq *Deque[T]) Len() int {
	return dq.len
}

