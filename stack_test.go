package goal

import (
	"testing"
)

func TestStackDataTypes(t *testing.T) {

	{
		s := MakeStack[bool](2)
		s.Push(true)
		s.Push(false)
		a_t, ok := s.Pop()
		if a_t || !ok {
			t.Errorf("bool error")
		}
		a_f, ok := s.Pop()
		if !a_f || !ok {
			t.Errorf("bool error")
		}
	}

	{
		s := MakeStack[uint](2)
		var a uint = 32
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("uint error")
		}
	}

	{
		s := MakeStack[int](2)
		var a int = 32
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("int error")
		}
	}

	{
		s := MakeStack[uint8](2)
		var a uint8 = 8
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("uint8 error")
		}
	}

	{
		s := MakeStack[uint16](2)
		var a uint16 = 16
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("uint16 error")
		}
	}

	{
		s := MakeStack[uint32](2)
		var a uint32 = 32
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("uint32 error")
		}
	}

	{
		s := MakeStack[uint64](2)
		var a uint64 = 64
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("uint64 error")
		}
	}

	{
		s := MakeStack[int8](2)
		var a int8 = 8
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("int8 error")
		}
	}

	{
		s := MakeStack[int16](2)
		var a int16 = 16
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("int16 error")
		}
	}

	{
		s := MakeStack[int32](2)
		var a int32 = 32
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("int32 error")
		}
	}

	{
		s := MakeStack[int64](2)
		var a int64 = 64
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("int64 error")
		}
	}

	{
		s := MakeStack[float32](2)
		var a float32 = 32.001
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("float32 error")
		}
	}

	{
		s := MakeStack[float64](2)
		var a float64 = 64.001
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("float64 error")
		}
	}

	{
		s := MakeStack[complex64](2)
		var a complex64 = 3 + 4i
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("complex64 error")
		}
	}

	{
		s := MakeStack[complex128](2)
		var a complex128 = 3 + 4i
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("complex128 error")
		}
	}

	{
		s := MakeStack[byte](2)
		var a byte = 2
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("byte error")
		}
	}

	{
		s := MakeStack[rune](2)
		var a rune = 2
		s.Push(a)
		ao, ok := s.Pop()
		if ao != a && !ok {
			t.Errorf("rune error")
		}
	}

	{
		s := MakeStack[[]byte](2)
		var a []byte = []byte{0, 1, 2, 3, 4}
		s.Push(a)
		ao, ok := s.Pop()
		for i, v := range ao {
			if v != a[i] && !ok {
				t.Errorf("byte slice error")
			}
		}
	}

	{
		s := MakeStack[*string](2)
		var a string = "test"
		s.Push(&a)
		ao, ok := s.Pop()
		if *ao != a && !ok {
			t.Errorf("ptr error")
		}
	}
}

func TestStackPushPop(t *testing.T) {

	s := MakeStack[byte](10)
	var a byte = 'a'
	var b byte = 'b'
	s.Push(a)
	s.Push(b)

	bo, ok := s.Pop()
	if bo != 'b' || !ok {
		t.Errorf("Pop error")
	}
	ao, ok := s.Pop()
	if ao != 'a' || !ok {
		t.Errorf("Pop error")
	}
}

func TestStackLenCap(t *testing.T) {

	s := MakeStack[int](10)
	var a int = 1
	s.Push(a)
	s.Push(a)
	if s.Len() != 2 {
		t.Errorf("Len() error")
	}
	if s.Cap() != 10 {
		t.Errorf("Cap() error")
	}
}

func TestStackGrow(t *testing.T) {

	s := MakeStack[int](10)

	if s.Len() != 0 {
		t.Errorf("new queue Len() error")
	}
	if s.Cap() != 10 {
		t.Errorf("new queue Cap() error")
	}

	var a int = 1
	for i := 0; i <= 10; i++ {
		s.Push(a)
	}

	if s.Len() != 11 {
		t.Errorf("grown queue Len() error")
	}
	if s.Cap() != 20 {
		t.Errorf("grown queue Cap() error")
	}
}
