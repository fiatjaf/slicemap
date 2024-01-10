package slicemap_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fiatjaf/slicemap"
)

func BenchmarkSliceMap(b *testing.B) {
	for _, n := range []int{10, 30, 90, 240} {
		b.Run(fmt.Sprintf("%d:store()", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := slicemap.New2[string, any]()
				for i := 0; i < n; i++ {
					m.Store(strconv.Itoa(i), i)
				}
			}
		})
		b.Run(fmt.Sprintf("%d:load()", n), func(b *testing.B) {
			m := slicemap.New2[string, any]()
			for i := 0; i < n; i++ {
				m.Store(strconv.Itoa(i), i)
			}
			for i := 0; i < b.N; i++ {
				_, _ = m.Load(strconv.Itoa(n * 2))
				_, _ = m.Load(strconv.Itoa(n / 2))
				_, _ = m.Load(strconv.Itoa(n * 2 / 3))
				_, _ = m.Load(strconv.Itoa(n - 1))
				_, _ = m.Load(strconv.Itoa(n / 4))
			}
		})
	}
}

func BenchmarkStdMap(b *testing.B) {
	for _, n := range []int{5, 10, 20, 40, 80, 160} {
		b.Run(fmt.Sprintf("%d:store()", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := make(map[string]any)
				for i := 0; i < n; i++ {
					m[strconv.Itoa(i)] = i
				}
			}
		})
		b.Run(fmt.Sprintf("%d:load()", n), func(b *testing.B) {
			m := make(map[string]any)
			for i := 0; i < n; i++ {
				m[strconv.Itoa(i)] = i
			}
			for i := 0; i < b.N; i++ {
				_, _ = m[strconv.Itoa(n*2)]
				_, _ = m[strconv.Itoa(n/2)]
				_, _ = m[strconv.Itoa(n*2/3)]
				_, _ = m[strconv.Itoa(n-1)]
				_, _ = m[strconv.Itoa(n/4)]
			}
		})
	}
}

func BenchmarkSliceMapInts(b *testing.B) {
	for _, n := range []int{5, 10, 20, 40, 80, 160} {
		b.Run(fmt.Sprintf("%d:store()", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := slicemap.New2[int, any]()
				for i := 0; i < n; i++ {
					m.Store(i, i)
				}
			}
		})
		b.Run(fmt.Sprintf("%d:load()", n), func(b *testing.B) {
			m := slicemap.New2[int, any]()
			for i := 0; i < n; i++ {
				m.Store(i, i)
			}
			for i := 0; i < b.N; i++ {
				_, _ = m.Load(n * 2)
				_, _ = m.Load(n / 2)
				_, _ = m.Load(n * 2 / 3)
				_, _ = m.Load(n - 1)
				_, _ = m.Load(n / 4)
			}
		})
	}
}

func BenchmarkStdMapInts(b *testing.B) {
	for _, n := range []int{5, 10, 20, 40, 80, 160} {
		b.Run(fmt.Sprintf("%d:store()", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m := make(map[int]any)
				for i := 0; i < n; i++ {
					m[i] = i
				}
			}
		})
		b.Run(fmt.Sprintf("%d:load()", n), func(b *testing.B) {
			m := make(map[int]any)
			for i := 0; i < n; i++ {
				m[i] = i
			}
			for i := 0; i < b.N; i++ {
				_, _ = m[n*2]
				_, _ = m[n/2]
				_, _ = m[n*2/3]
				_, _ = m[n-1]
				_, _ = m[n/4]
			}
		})
	}
}
