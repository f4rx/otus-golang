package hw04lrucache

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())
		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}

func TestList2(t *testing.T) {
	lpb := List.PushBack
	lpf := List.PushFront
	lr := List.Remove

	type action struct {
		name   string
		run    func(l List)
		result []int
	}

	actions := []action{
		{
			name: "push back and front",
			run: func(l List) {
				lpb(l, 10)
				lpf(l, 5)
				lpb(l, 20)
				lpf(l, 1)
			},
			result: []int{1, 5, 10, 20},
		},
		{
			name: "remove last item",
			run: func(l List) {
				lpf(l, 10)
				lpf(l, 20)
				i := lpf(l, 20)
				lr(l, i)
			},
			result: []int{20, 10},
		},
		{
			name: "remove midle item",
			run: func(l List) {
				lpf(l, 10)
				i := lpf(l, 20)
				lpf(l, 30)
				lr(l, i)
			},
			result: []int{30, 10},
		},
		{
			name: "remove first item",
			run: func(l List) {
				i := lpf(l, 10)
				lpf(l, 20)
				lpf(l, 30)
				lr(l, i)
			},
			result: []int{30, 20},
		},
		{
			name: "move to front single item",
			run: func(l List) {
				l.PushFront(10)
				l.MoveToFront(l.Front())
			},
			result: []int{10},
		},
		{
			name: "remove single item",
			run: func(l List) {
				i := l.PushFront(10)
				l.Remove(i)
				require.Equal(t, 0, l.Len())
				require.Nil(t, l.Front())
				require.Nil(t, l.Back())
			},
			result: []int{},
		},
	}

	for i, a := range actions {
		l := NewList()
		t.Run(fmt.Sprintf("test_%d, name:%s", i, a.name), func(t *testing.T) {
			a.run(l)
			slog.Debug(l)
			elems := make([]int, 0, l.Len())
			for i := l.Front(); i != nil; i = i.Next {
				elems = append(elems, i.Value.(int))
			}
			require.Equal(t, a.result, elems)
		})
	}
}
