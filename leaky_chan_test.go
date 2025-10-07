package goal

import (
	"testing"
	"time"
)

// go test -v
// go test -bench=.

func NewLeakyChan[T any](buffer int) LeakyChan[T] {
	return make(chan T, buffer)
}

// TestLeakyChan_BasicOperations
func TestLeakyChan_BasicOperations(t *testing.T) {
	ch := NewLeakyChan[int](1)

	// Тест отправки и получения
	ch <- 42
	select {
	case val := <-ch:
		if val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}
	case <-time.After(time.Second):
		t.Error("Timeout while receiving from channel")
	}
}

// TestLeakyChan_Concurrent
func TestLeakyChan_Concurrent(t *testing.T) {
	ch := NewLeakyChan[string](10)
	const numMessages = 100

	go func() {
		for i := 0; i < numMessages; i++ {
			ch <- "message"
		}
	}()

	go func() {
		for i := 0; i < numMessages; i++ {
			select {
			case <-ch:
				// Success
			case <-time.After(time.Second):
				t.Error("Timeout while receiving message")
				return
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)
}

// TestLeakyChan_Close
func TestLeakyChan_Close(t *testing.T) {
	ch := NewLeakyChan[int](0)

	go func() {
		ch <- 1
		close(ch)
	}()

	val, ok := <-ch
	if !ok || val != 1 {
		t.Errorf("Expected 1, true, got %d, %t", val, ok)
	}

	_, ok = <-ch
	if ok {
		t.Error("Channel should be closed")
	}
}

// TestLeakyChan_Buffer
func TestLeakyChan_Buffer(t *testing.T) {
	bufferSize := 3
	ch := NewLeakyChan[int](bufferSize)

	for i := 0; i < bufferSize; i++ {
		ch <- i
	}

	for i := 0; i < bufferSize; i++ {
		select {
		case val := <-ch:
			if val != i {
				t.Errorf("Expected %d, got %d", i, val)
			}
		case <-time.After(time.Second):
			t.Errorf("Timeout while receiving message %d", i)
		}
	}
}

// TestLeakyChan_WithStruct
func TestLeakyChan_WithStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	ch := NewLeakyChan[Person](1)

	person := Person{Name: "Alice", Age: 30}
	ch <- person

	select {
	case received := <-ch:
		if received.Name != person.Name || received.Age != person.Age {
			t.Errorf("Expected %+v, got %+v", person, received)
		}
	case <-time.After(time.Second):
		t.Error("Timeout while receiving struct")
	}
}

// TestLeakyChan_Blocking
func TestLeakyChan_Blocking(t *testing.T) {
	ch := NewLeakyChan[int](0)

	done := make(chan bool)

	go func() {
		<-ch
		done <- true
	}()

	time.Sleep(50 * time.Millisecond)

	ch <- 1

	select {
	case <-done:
		// Success
	case <-time.After(time.Second):
		t.Error("Receiver should have been unblocked")
	}
}

func TestLeakyChan_BasicOperationsWithFirstLeaks(t *testing.T) {
	ch := NewLeakyChan[int](3)

	values := []int{10, 20, 30, 40}
	for _, value := range values {
		ch.SendLeakFirst(value)
	}

	expected := []int{20, 30, 40}

	for i, value := range expected {
		select {
		case received := <-ch:
			if received != value {
				t.Errorf("Value %d: expected %d, got %d", i, value, received)
			}
		case <-time.After(time.Second):
			t.Errorf("Timeout while receiving value %d", i)
		}
	}

	select {
	case val := <-ch:
		t.Errorf("Channel should be empty, but received: %d", val)
	default:

	}
}

func TestLeakyChan_BasicOperationsWithLastLeaks(t *testing.T) {
	ch := NewLeakyChan[int](3)

	values := []int{10, 20, 30, 40}
	for _, value := range values {
		ch.SendLeakLast(value)
	}

	expected := []int{10, 20, 30}

	for i, value := range expected {
		select {
		case received := <-ch:
			if received != value {
				t.Errorf("Value %d: expected %d, got %d", i, value, received)
			}
		case <-time.After(time.Second):
			t.Errorf("Timeout while receiving value %d", i)
		}
	}

	select {
	case val := <-ch:
		t.Errorf("Channel should be empty, but received: %d", val)
	default:

	}
}
