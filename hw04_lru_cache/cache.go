package hw04lrucache

import (
	"fmt"
	"strings"
	"sync"
)

var mutex = &sync.Mutex{}

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type cacheItem struct {
	// key   string
	key   Key
	value interface{}
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	mutex.Lock()
	defer mutex.Unlock()

	// slog.Debug("Set: ", key, " : ", value)
	if item, ok := l.items[key]; ok {
		// slog.Debug("Update cache ", item)
		currentCacheItem := item.Value.(cacheItem)
		currentCacheItem.value = value
		item.Value = currentCacheItem
		l.queue.MoveToFront(item)
		// slog.Debug(l.queue.Len())
		// slog.Debug(l.queue)
		return true
	} else { //nolint:golint
		if l.queue.Len() == l.capacity {
			item = l.queue.Back()
			backCacheItem := item.Value.(cacheItem)
			l.queue.Remove(item)
			delete(l.items, backCacheItem.key)
		}
		newCacheItem := cacheItem{key, value}
		item = l.queue.PushFront(newCacheItem)
		l.items[key] = item
		// slog.Debug(l.queue.Len())
		// slog.Debug(l.queue)
	}
	return false
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if item, ok := l.items[key]; ok {
		currentCacheItem := item.Value.(cacheItem)
		l.queue.MoveToFront(item)
		l.items[key] = l.queue.Front()
		return currentCacheItem.value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	mutex.Lock()
	defer mutex.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func (l *lruCache) String() string {
	var out strings.Builder
	out.WriteString(fmt.Sprintf("%s\n", l.queue))
	// fmt.Println(l.queue)
	out.WriteString("Keys: ")
	for k := range l.items {
		out.WriteString(fmt.Sprintf("%s ", k))
	}

	return out.String()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
