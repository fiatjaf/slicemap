package slicemap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type Map[K constraints.Ordered, V any] struct {
	keys   []K
	values []V
}

func New[K constraints.Ordered, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

func (m *Map[K, V]) Clear() {
	m.keys = nil
	m.values = nil
}

func (m *Map[K, V]) Delete(key K) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		m.keys = slices.Delete(m.keys, idx, idx+1)
		m.values = slices.Delete(m.values, idx, idx+1)
	}
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		return m.values[idx], true
	}
	return value, false
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		value = m.values[idx]
		m.keys = slices.Delete(m.keys, idx, idx+1)
		m.values = slices.Delete(m.values, idx, idx+1)
		return value, true
	}
	return value, false
}

func (m *Map[K, V]) LoadAndStore(key K, value V) (actual V, loaded bool) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		actual = m.values[idx]
		m.values[idx] = value
		return actual, true
	} else {
		m.keys = append(m.keys, key) // just to increase the slice capacity
		copy(m.keys[idx+1:], m.keys[idx:])
		m.keys[idx] = key
		m.values = append(m.values, value) // just to increase the slice capacity
		copy(m.values[idx+1:], m.values[idx:])
		m.values[idx] = value

		actual = value
		return actual, false
	}
}

func (m *Map[K, V]) LoadOrCompute(key K, valueFn func() V) (actual V, loaded bool) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		actual = m.values[idx]
		return actual, true
	} else {
		m.keys = append(m.keys, key) // just to increase the slice capacity
		copy(m.keys[idx+1:], m.keys[idx:])
		m.keys[idx] = key

		value := valueFn()
		m.values = append(m.values, value) // just to increase the slice capacity
		copy(m.values[idx+1:], m.values[idx:])
		m.values[idx] = value

		return value, false
	}
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		actual = m.values[idx]
		return actual, true
	} else {
		m.keys = append(m.keys, key) // just to increase the slice capacity
		copy(m.keys[idx+1:], m.keys[idx:])
		m.keys[idx] = key

		m.values = append(m.values, value) // just to increase the slice capacity
		copy(m.values[idx+1:], m.values[idx:])
		m.values[idx] = value

		return value, false
	}
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	for i := range m.keys {
		if !f(m.keys[i], m.values[i]) {
			break
		}
	}
}

func (m *Map[K, V]) Size() int {
	return len(m.keys)
}

func (m *Map[K, V]) Store(key K, value V) {
	if idx, ok := slices.BinarySearch(m.keys, key); ok {
		m.values[idx] = value
	} else {
		m.keys = append(m.keys, key) // just to increase the slice capacity
		copy(m.keys[idx+1:], m.keys[idx:])
		m.keys[idx] = key

		m.values = append(m.values, value) // just to increase the slice capacity
		copy(m.values[idx+1:], m.values[idx:])
		m.values[idx] = value
	}
}
