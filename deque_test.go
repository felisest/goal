package goal

import (
	"testing"
)

func TestDequeDataTypes(t *testing.T) {

	{
		dq := MakeDeque[bool]()
		dq.PushBack(true)
		dq.PushBack(false)
		a_t, ok := dq.PopBack()
		if a_t || !ok {
			t.Errorf("bool error")
		}
		a_f, ok := dq.PopBack()
		if !a_f || !ok {
			t.Errorf("bool error")
		}
	}

	{
		dq := MakeDeque[uint]()
		var a uint = 32
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("uint error")
		}
	}

	{
		dq := MakeDeque[int]()
		var a int = 32
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("int error")
		}
	}

	{
		dq := MakeDeque[uint8]()
		var a uint8 = 8
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("uint8 error")
		}
	}

	{
		dq := MakeDeque[uint16]()
		var a uint16 = 16
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("uint16 error")
		}
	}

	{
		dq := MakeDeque[uint32]()
		var a uint32 = 32
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("uint32 error")
		}
	}

	{
		dq := MakeDeque[uint64]()
		var a uint64 = 64
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("uint64 error")
		}
	}

	{
		dq := MakeDeque[int8]()
		var a int8 = 8
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("int8 error")
		}
	}

	{
		dq := MakeDeque[int16]()
		var a int16 = 16
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("int16 error")
		}
	}

	{
		dq := MakeDeque[int32]()
		var a int32 = 32
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("int32 error")
		}
	}

	{
		dq := MakeDeque[int64]()
		var a int64 = 64
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("int64 error")
		}
	}

	{
		dq := MakeDeque[float32]()
		var a float32 = 32.001
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("float32 error")
		}
	}

	{
		dq := MakeDeque[float64]()
		var a float64 = 64.001
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("float64 error")
		}
	}

	{
		dq := MakeDeque[complex64]()
		var a complex64 = 3 + 4i
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("complex64 error")
		}
	}

	{
		dq := MakeDeque[complex128]()
		var a complex128 = 3 + 4i
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("complex128 error")
		}
	}

	{
		dq := MakeDeque[byte]()
		var a byte = 2
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("byte error")
		}
	}

	{
		dq := MakeDeque[rune]()
		var a rune = 2
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		if ao != a && !ok {
			t.Errorf("rune error")
		}
	}

	{
		dq := MakeDeque[[]byte]()
		var a []byte = []byte{0, 1, 2, 3, 4}
		dq.PushBack(a)
		ao, ok := dq.PopBack()
		for i, v := range ao {
			if v != a[i] && !ok {
				t.Errorf("byte slice error")
			}
		}
	}

	{
		dq := MakeDeque[*string]()
		var a string = "test"
		dq.PushBack(&a)
		ao, ok := dq.PopBack()
		if *ao != a && !ok {
			t.Errorf("ptr error")
		}
	}
}

func TestDequePushPopBack(t *testing.T) {

	d := MakeDeque[byte]()
	var a byte = 'a'
	var b byte = 'b'
	d.PushBack(a)
	d.PushBack(b)

	va, ok := d.PopBack()
	if va != 'b' || !ok {
		t.Errorf("PopBack error")
	}
	vb, ok := d.PopBack()
	if vb != 'a' || !ok {
		t.Errorf("PopBack error")
	}
}

func TestDequePushPopFront(t *testing.T) {

	d := MakeDeque[byte]()
	var a byte = 'a'
	var b byte = 'b'
	d.PushFront(a)
	d.PushFront(b)

	va, ok := d.PopFront()
	if va != 'b' || !ok {
		t.Errorf("PopFront error")
	}
	vb, ok := d.PopFront()
	if vb != 'a' || !ok {
		t.Errorf("PopFront error")
	}
}

func TestDequePushPopCombined(t *testing.T) {

	d := MakeDeque[byte]()
	var a byte = 'a'
	var b byte = 'b'
	d.PushFront(a)
	d.PushFront(b)

	va, ok := d.PopBack()
	if va != 'a' || !ok {
		t.Errorf("Pop error")
	}
	vb, ok := d.PopBack()
	if vb != 'b' || !ok {
		t.Errorf("Pop error")
	}
}