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

func TestList1(t *testing.T) {
	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushFront(20) // [10]
		l.PushFront(30) // [10]
		l.PushBack(100) // [10, 20]
		l.PushBack(110) // [10, 20]
		l.PushFront(5)  // [10]
		// l.PushBack(30)  // [10, 20, 30]

		l.MoveToFront(l.Front())
		// l.MoveToFront(l.Back())

		slog.Debug(l)

		for i := l.Front(); i != nil; i = i.Next {
			slog.Debug(i.Value.(int))
		}
	})
}

func TestListMoveToFrontSignleItem(t *testing.T) {
	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)
		l.MoveToFront(l.Front())
		slog.Debug(l)

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{10}, elems)
	})
}

func TestListRemoveSingleItem(t *testing.T) {
	t.Run("complex", func(t *testing.T) {
		l := NewList()

		i := l.PushFront(10) // [10]
		l.Remove(i)

		slog.Debug(l)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
}

// func TestListRemoveFirstItem(t *testing.T) {
// 	t.Run("complex", func(t *testing.T) {
// 		l := NewList()

// 		i := l.PushFront(10) // [10]
// 		l.PushFront(20)      // [10]
// 		l.PushFront(30)      // [10]

// 		l.Remove(i)

// 		slog.Debug(l)
// 		require.Equal(t, 2, l.Len())

// 		elems := make([]int, 0, l.Len())
// 		for i := l.Front(); i != nil; i = i.Next {
// 			elems = append(elems, i.Value.(int))
// 		}
// 		require.Equal(t, []int{30, 20}, elems)
// 	})
// }

// func TestListRemoveMidleItem(t *testing.T) {
// 	t.Run("complex", func(t *testing.T) {
// 		l := NewList()

// 		l.PushFront(10)
// 		i := l.PushFront(20)
// 		l.PushFront(30)
// 		l.Remove(i)

// 		slog.Debug(l)
// 		require.Equal(t, 2, l.Len())

// 		elems := make([]int, 0, l.Len())
// 		for i := l.Front(); i != nil; i = i.Next {
// 			elems = append(elems, i.Value.(int))
// 		}
// 		require.Equal(t, []int{30, 10}, elems)
// 	})
// }

// func TestListRemoveLastItem(t *testing.T) {
// 	t.Run("complex", func(t *testing.T) {
// 		l := NewList()

// 		l.PushFront(10)      // [10]
// 		l.PushFront(20)      // [10]
// 		i := l.PushFront(30) // [10]

// 		l.Remove(i)

// 		slog.Debug(l)
// 		elems := make([]int, 0, l.Len())
// 		for i := l.Front(); i != nil; i = i.Next {
// 			elems = append(elems, i.Value.(int))
// 		}
// 		require.Equal(t, []int{20, 10}, elems)
// 	})
// }

// func TestListPushBackInEmpty(t *testing.T) {
// 	t.Run("complex", func(t *testing.T) {
// 		l := NewList()

// 		l.PushBack(10) // [10]
// 		l.PushFront(5) // [5 ,10 ]
// 		l.PushBack(20) // [5 ,10, 20 ]
// 		l.PushFront(1) // [1, 5 ,10, 20 ]

// 		slog.Debug(l)
// 		elems := make([]int, 0, l.Len())
// 		for i := l.Front(); i != nil; i = i.Next {
// 			elems = append(elems, i.Value.(int))
// 		}
// 		require.Equal(t, []int{1, 5, 10, 20}, elems)
// 	})
// }

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
		action{
			name: "push back and front",
			run: func(l List) {
				lpb(l, 10)
				lpf(l, 5)
				lpb(l, 20)
				lpf(l, 1)
			},
			result: []int{1, 5, 10, 20},
		},
		action{
			name: "remove last item",
			run: func(l List) {
				lpf(l, 10)
				lpf(l, 20)
				i := lpf(l, 20)
				lr(l, i)
			},
			result: []int{20, 10},
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

	// testActionsAndResult1 := func(l List) (string, []func(), []int) {
	// 	name := "push back and front"
	// 	funcList := []func(){
	// 		func() { lpb(l, 10) },
	// 		func() { lpf(l, 5) },
	// 		func() { lpb(l, 20) },
	// 		func() { lpf(l, 1) },
	// 	}
	// 	result := []int{1, 5, 10, 20}
	// 	return name, funcList, result
	// }

	// testActionsAndResult2 := func(l List) (string, []func(), []int) {
	// 	name := "remove last item"
	// 	funcList := []func(){
	// 		func() { lpf(l, 10) },
	// 		func() { lpf(l, 20) },
	// 		func() {
	// 			i := lpf(l, 20)
	// 			lr(l, i)
	// 		},
	// 	}
	// 	result := []int{20, 10}
	// 	return name, funcList, result
	// }

	// testActionsAndResult3 := func(l List) (string, []func(), []int) {
	// 	name := "remove midle item"
	// 	funcList := []func(){
	// 		func() { lpf(l, 10) },
	// 		func() {
	// 			i := lpf(l, 20)
	// 			lpf(l, 30)
	// 			lr(l, i)
	// 		},
	// 	}
	// 	result := []int{30, 10}
	// 	return name, funcList, result
	// }

	// testActionsAndResult4 := func(l List) (string, []func(), []int) {
	// 	name := "remove first item"
	// 	funcList := []func(){
	// 		func() {
	// 			i := lpf(l, 10)
	// 			lpf(l, 20)
	// 			lpf(l, 30)
	// 			lr(l, i)
	// 		},
	// 	}
	// 	result := []int{30, 20}
	// 	return name, funcList, result
	// }

	// testActionsAndResultList := []func(l List) (string, []func(), []int){
	// 	testActionsAndResult1,
	// 	testActionsAndResult2,
	// 	testActionsAndResult3,
	// 	testActionsAndResult4,
	// }

	// for i, testActionsAndResult := range testActionsAndResultList {
	// 	l := NewList()
	// 	name, actions, result := testActionsAndResult(l)
	// 	t.Run(fmt.Sprintf("test_%d, name:%s", i, name), func(t *testing.T) {
	// 		for _, action := range actions {
	// 			action()
	// 		}
	// 		slog.Debug(l)
	// 		elems := make([]int, 0, l.Len())
	// 		for i := l.Front(); i != nil; i = i.Next {
	// 			elems = append(elems, i.Value.(int))
	// 		}
	// 		require.Equal(t, result, elems)
	// 	})
	// }
}
