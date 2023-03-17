package syncPool

import "sync"

type Pool[T interface{}] struct {
	pool sync.Pool
	new  func() T
}

func NewPool[T interface{}](newFunc func() any) *Pool[T] {
	return &Pool[T]{
		pool: sync.Pool{New: newFunc},
	}
}

func (p *Pool[T]) Put(v T) {
	p.pool.Put(v)
}

func (p *Pool[T]) Get() T {
	return p.pool.Get().(T)
}
