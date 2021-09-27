package hw04lrucache

import (
	"math/rand"
	"sort"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		// Write me
	})
}

func TestCacheMultithreading(t *testing.T) {
	// t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

func TestCacheDummy(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(2)

		c.Set("aaa", 100)
		c.Set("bbb", 200)
		c.Get("aaa")
		// c.Get("bbb")
		c.Set("aaa", 300)
		slog.Debug(c)
	})
}

func TestCacheOverflow(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)
		c.Set("d", 400)
		// c.Get("bbb")
		slog.Debug(c)
		lruC := c.(*lruCache)
		slog.Debug(lruC.items)

		keys := make([]string, 0, 3)
		for k := range lruC.items {
			keys = append(keys, string(k))
		}
		sort.Strings(keys)
		require.Equal(t, []string{"b", "c", "d"}, keys)

		elems := make([]int, 0, lruC.queue.Len())
		for i := lruC.queue.Front(); i != nil; i = i.Next {
			cacheItem := i.Value.(cacheItem)
			value := cacheItem.value.(int)
			elems = append(elems, value)
		}
		require.Equal(t, []int{400, 300, 200}, elems)
	})
}

func TestCacheComplexOverflow(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Set("c", 300)
		c.Set("a", 200)
		c.Get("b")
		c.Get("c")
		c.Set("d", 400)
		c.Set("e", 500)
		slog.Debug(c)

		lruC := c.(*lruCache)
		slog.Debug(lruC.items)

		keys := make([]string, 0, 3)
		for k := range lruC.items {
			keys = append(keys, string(k))
		}
		sort.Strings(keys)
		require.Equal(t, []string{"c", "d", "e"}, keys)

		elems := make([]int, 0, lruC.queue.Len())
		for i := lruC.queue.Front(); i != nil; i = i.Next {
			cacheItem := i.Value.(cacheItem)
			value := cacheItem.value.(int)
			elems = append(elems, value)
		}
		require.Equal(t, []int{500, 400, 300}, elems)
	})
}

func TestClear(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(2)

		c.Set("a", 100)
		c.Set("b", 200)
		c.Clear()
		slog.Debug(c)

		lruC := c.(*lruCache)

		require.Equal(t, 0, lruC.queue.Len())

		// Проверяем что после очистки всё дальше работает
		c.Set("a", 100)
		c.Set("b", 200)
		slog.Debug(lruC)
		require.Equal(t, 2, lruC.queue.Len())
	})
}
