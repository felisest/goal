package goal

import (
	"testing"
)

func TestQueueDataTypes(t *testing.T) {

	{
		q := MakeQueue[bool](2)
		q.Push(true)
		q.Push(false)
		a_t, ok := q.Pop()
		if !a_t || !ok {
			t.Errorf("bool error")
		}
		a_f, ok := q.Pop()
		if a_f || !ok {
			t.Errorf("bool error")
		}
	}

	{
		q := MakeQueue[uint](2)
		var a uint = 32
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("uint error")
		}
	}

	{
		q := MakeQueue[int](2)
		var a int = 32
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("int error")
		}
	}

	{
		q := MakeQueue[uint8](2)
		var a uint8 = 8
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("uint8 error")
		}
	}

	{
		q := MakeQueue[uint16](2)
		var a uint16 = 16
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("uint16 error")
		}
	}

	{
		q := MakeQueue[uint32](2)
		var a uint32 = 32
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("uint32 error")
		}
	}

	{
		q := MakeQueue[uint64](2)
		var a uint64 = 64
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("uint64 error")
		}
	}

	{
		q := MakeQueue[int8](2)
		var a int8 = 8
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("int8 error")
		}
	}

	{
		q := MakeQueue[int16](2)
		var a int16 = 16
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("int16 error")
		}
	}

	{
		q := MakeQueue[int32](2)
		var a int32 = 32
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("int32 error")
		}
	}

	{
		q := MakeQueue[int64](2)
		var a int64 = 64
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("int64 error")
		}
	}

	{
		q := MakeQueue[float32](2)
		var a float32 = 32.001
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("float32 error")
		}
	}

	{
		q := MakeQueue[float64](2)
		var a float64 = 64.001
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("float64 error")
		}
	}

	{
		q := MakeQueue[complex64](2)
		var a complex64 = 3 + 4i
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("complex64 error")
		}
	}

	{
		q := MakeQueue[complex128](2)
		var a complex128 = 3 + 4i
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("complex128 error")
		}
	}

	{
		q := MakeQueue[byte](2)
		var a byte = 2
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("byte error")
		}
	}

	{
		q := MakeQueue[rune](2)
		var a rune = 2
		q.Push(a)
		ao, ok := q.Pop()
		if ao != a && !ok {
			t.Errorf("rune error")
		}
	}

	{
		q := MakeQueue[[]byte](2)
		var a []byte = []byte{0, 1, 2, 3, 4}
		q.Push(a)
		ao, ok := q.Pop()
		for i, v := range ao {
			if v != a[i] && !ok {
				t.Errorf("byte slice error")
			}
		}
	}

	{
		q := MakeQueue[*string](2)
		var a string = "test"
		q.Push(&a)
		ao, ok := q.Pop()
		if *ao != a && !ok {
			t.Errorf("ptr error")
		}
	}
}

func TestQueueLenCap(t *testing.T) {

	q := MakeQueue[int](10)
	var a int = 1
	q.Push(a)
	q.Push(a)
	if q.Len() != 2 {
		t.Errorf("Len() error")
	}
	if q.Cap() != 10 {
		t.Errorf("Cap() error")
	}
}

func TestQueueGrow(t *testing.T) {
	
	q := MakeQueue[int](10)

	if q.Len() != 0 {
		t.Errorf("new queue Len() error")
	}
	if q.Cap() != 10 {
		t.Errorf("new queue Cap() error")
	}

	var a int = 1
	for i := 0; i <= 10; i++ {
		q.Push(a)
	}

	if q.Len() != 11 {
		t.Errorf("grown queue Len() error")
	}
	if q.Cap() != 20 {
		t.Errorf("grown queue Cap() error")
	}
}