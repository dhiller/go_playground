package counter

import "sync"

type Counter interface {
	Inc()
	Dec()
	Value() int
}

type ThreadSafeCounter struct {
	value int
	mutex sync.Mutex
}

func (counter *ThreadSafeCounter) Inc() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value++
}

func (counter *ThreadSafeCounter) Dec() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value--
}

func (counter *ThreadSafeCounter) Value() int {
	return counter.value
}
