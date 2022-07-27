package goal

// Queue represents a instance of the queue data container.
type Queue[T any] struct {
	buff  []T
	begin int
	end   int
	len   int
	cap   int
}

func grow_container(cur_cap int) int {

	mnum := 512
	i := 0
	for ; (cur_cap > mnum<<i) && (i < 4); i++ {
	}

	return cur_cap + cur_cap/(1<<i)
}

// MakeQueue constructs and return a new Queue.
func MakeQueue[T any](capacity int) *Queue[T] {

	return &Queue[T]{buff: make([]T, capacity), cap: capacity}
}

func (q *Queue[T]) resize() {

	new_capacity := grow_container(q.cap)
	new_buff := make([]T, new_capacity)

	if q.begin >= q.end {
		num := copy(new_buff, q.buff[q.begin:])
		copy(new_buff[num:], q.buff[:q.end])
	} else {
		copy(new_buff, q.buff[:q.end])
	}

	q.buff = new_buff
	q.begin = 0
	q.end = q.len
	q.cap = new_capacity
}

// Push inserts element at the end.
func (q *Queue[T]) Push(v T) {

	if q.len >= len(q.buff) {
		q.resize()
	}

	q.buff[q.end] = v

	if q.end >= len(q.buff)-1 && q.begin > 0 {
		q.end = 0
	} else {
		q.end++
	}

	q.len++
}

// Pop return the first element.
func (q *Queue[T]) Pop() (T, bool) {

	var dr T
	if q.len == 0 {
		return dr, false
	}

	v := q.buff[q.begin]
	q.buff[q.begin] = dr

	if q.begin+1 > len(q.buff) {
		q.begin = 0
	} else {
		q.begin++
	}

	q.len--

	return v, true
}

// Erase erases all unused elements.
func (q *Queue[T]) Erase() {
	new_buff := make([]T, q.cap)

	if q.begin >= q.end {
		num := copy(new_buff, q.buff[q.begin:])
		copy(new_buff[num:], q.buff[:q.end])
	} else {
		copy(new_buff, q.buff[:q.end])
	}

	q.buff = new_buff
	q.begin = 0
	q.end = q.len
}

// Clear clears the container. After this call, Len() returns zero.
func (q *Queue[T]) Clear() {
	new_buff := make([]T, q.cap)
	q.buff = new_buff
}

// Len returns the number of elements.
func (q *Queue[T]) Len() int {
	return q.len
}

// Cap returns the capacity of container.
func (q *Queue[T]) Cap() int {
	return q.cap
}
