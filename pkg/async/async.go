package async

import "sync"

type Task[T any] struct {
	result chan T
	wg     sync.WaitGroup
}

func Async[T any](f func() T) *Task[T] {
	t := &Task[T]{result: make(chan T)}
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		t.result <- f()
	}()

	return t
}

func (t *Task[T]) Await() T {
	go func() {
		t.wg.Wait()
		close(t.result)
	}()

	return <-t.result
}
