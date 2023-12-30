package slicemap_test

import (
	"testing"

	"github.com/fiatjaf/slicemap"
)

func TestMap(t *testing.T) {
	m := slicemap.New[string, any]()
	m.Store("a", 4654)
	m.Store("a", 1)
	m.Store("b", 1515)
	m.Store("z", 2)
	if v, _ := m.LoadAndDelete("z"); v != 2 {
		t.Fatalf("LoadAndDelete failed")
	}
	m.LoadOrStore("b", 7)
	if v, ok := m.LoadAndStore("b", 2); v != 1515 || !ok {
		t.Fatalf("LoadAndStore failed")
	}
	if v, _ := m.Load("b"); v != 2 {
		t.Fatalf("Load failed")
	}
	if v, ok := m.LoadOrCompute("c", func() any { return 3 }); v != 3 || ok {
		t.Fatalf("LoadOrCompute failed")
	}
	count := 0
	m.Range(func(key string, value any) bool {
		count++
		return true
	})
	if count != 3 || count != m.Size() {
		t.Fatalf("Size or Range failed")
	}
	if v, ok := m.Load("a"); v != 1 || !ok {
		t.Fatalf("Load failed")
	}
	m.Clear()
	if m.Size() != 0 {
		t.Fatalf("Clear failed")
	}
}
