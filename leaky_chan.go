package goal

type LeakyChan[T any] chan T

func (c *LeakyChan[T]) SendLeakLast(value T) bool {
	select {
	case *c <- value:
		return true
	default:
		return false
	}
}

func (c *LeakyChan[T]) SendLeakFirst(value T) bool {
	select {
	case *c <- value:
		return true
	default:
		<-*c
		c.SendLeakFirst(value)
		return false
	}
}
