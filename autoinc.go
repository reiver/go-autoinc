package autoinc

import (
	"sync"
)

type incrementable interface {
	uint8 | uint16 | uint32 | uint64
}

// AutoInc is an implementation of an AUTO_INCREMENT type.
//
// AutoInc is thread-safe.
// I.e., AutoInc is safe to use with go-routines.
type AutoInc[T incrementable] struct {
	mutex sync.Mutex
	current T
}

// Current returns the current AUTO_INCREMENT value.
//
// If you are not sure whether to to use Current or to use [Next], use [Next].
func (receiver *AutoInc[T]) Current() T {
	if nil == receiver {
		var nada T
		return nada
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	return receiver.current
}

// Next returns the next AUTO_INCREMENT value.
//
// If you are not sure whether to to use Next or to use [Current], use Next.
func (receiver *AutoInc[T]) Next() (T, error) {
	if nil == receiver {
		var nada T
		return nada, ErrNilReceiver
	}

	receiver.mutex.Lock()
	defer receiver.mutex.Unlock()

	var previous T = receiver.current

	receiver.current++
	var current T = receiver.current

	if current < previous {
		return current, ErrOverFlow
	}

	return current, nil
}
