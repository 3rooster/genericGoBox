package syncMap

import "sync"

type Map[KT interface{}, VT interface{}] struct {
	store sync.Map
}

func (m *Map[KT, VT]) Load(key KT) (v VT, exist bool) {
	res, o := m.store.Load(key)
	if !o {
		return v, o
	}
	return res.(VT), o
}

func (m *Map[KT, VT]) Store(key KT, v VT) {
	m.store.Store(key, v)
}

func (m *Map[KT, VT]) LoadOrStore(key KT, v VT) (actual VT, loaded bool) {
	res, ok := m.store.LoadOrStore(key, v)
	if res == nil {
		return actual, ok
	}
	return res.(VT), ok
}

func (m *Map[KT, VT]) Delete(key KT) {
	m.store.Delete(key)
}

func (m *Map[KT, VT]) Range(rangeFunc func(k KT, v VT) bool) {
	m.store.Range(func(key, value any) bool {
		return rangeFunc(key.(KT), value.(VT))
	})
}
