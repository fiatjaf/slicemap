package slicemap

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type entry[K constraints.Ordered, V any] struct {
	key   K
	value V
}

func entryCompare[K constraints.Ordered, V any](a entry[K, V], b K) int {
	if a.key > b {
		return 1
	}
	if b > a.key {
		return -1
	}
	return 0
}

type Map2[K constraints.Ordered, V any] struct {
	entries []entry[K, V]
}

func New2[K constraints.Ordered, V any]() *Map[K, V] {
	return &Map[K, V]{}
}

func (m *Map2[K, V]) Clear() {
	m.entries = nil
}

func (m *Map2[K, V]) Delete(key K) {
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		m.entries = slices.Delete(m.entries, idx, idx+1)
	}
}

func (m *Map2[K, V]) Load(key K) (value V, ok bool) {
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		return m.entries[idx].value, true
	}
	return value, false
}

func (m *Map2[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		value = m.entries[idx].value
		m.entries = slices.Delete(m.entries, idx, idx+1)
		return value, true
	}
	return value, false
}

func (m *Map2[K, V]) LoadAndStore(key K, value V) (actual V, loaded bool) {
	entry := entry[K, V]{key, value}
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		actual = m.entries[idx].value
		m.entries[idx] = entry
		return actual, true
	} else {
		m.entries = append(m.entries, entry) // just to increase the slice capacity
		copy(m.entries[idx+1:], m.entries[idx:])
		m.entries[idx] = entry

		actual = value
		return actual, false
	}
}

func (m *Map2[K, V]) LoadOrCompute(key K, valueFn func() V) (actual V, loaded bool) {
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		actual = m.entries[idx].value
		return actual, true
	} else {
		value := valueFn()
		entry := entry[K, V]{key, value}
		m.entries = append(m.entries, entry) // just to increase the slice capacity
		copy(m.entries[idx+1:], m.entries[idx:])
		m.entries[idx] = entry
		return value, false
	}
}

func (m *Map2[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	entry := entry[K, V]{key, value}
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		actual = m.entries[idx].value
		return actual, true
	} else {
		m.entries = append(m.entries, entry) // just to increase the slice capacity
		copy(m.entries[idx+1:], m.entries[idx:])
		m.entries[idx] = entry

		return value, false
	}
}

func (m *Map2[K, V]) Range(f func(key K, value V) bool) {
	for i := range m.entries {
		if !f(m.entries[i].key, m.entries[i].value) {
			break
		}
	}
}

func (m *Map2[K, V]) Size() int {
	return len(m.entries)
}

func (m *Map2[K, V]) Store(key K, value V) {
	entry := entry[K, V]{key, value}
	if idx, ok := slices.BinarySearchFunc(m.entries, key, entryCompare); ok {
		m.entries[idx] = entry
	} else {
		m.entries = append(m.entries, entry) // just to increase the slice capacity
		copy(m.entries[idx+1:], m.entries[idx:])
		m.entries[idx] = entry
	}
}
