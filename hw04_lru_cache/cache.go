package hw04lrucache

import "sync"

type Key string

type cacheItem struct {
	key   Key
	value interface{}
}

type LruCache interface {
	sync.Locker

	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	sync.Mutex

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	val, ok := c.items[key]

	if ok {
		c.queue.MoveToFront(val)

		preparedVal, ok := val.Value.(*cacheItem)

		if !ok {
			return nil, false
		}

		return preparedVal.value, true
	}

	return nil, false
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	ci := &cacheItem{
		key:   key,
		value: value,
	}

	if listItem, ok := c.items[key]; ok {
		c.items[key].Value = ci
		c.queue.MoveToFront(listItem)
		return true
	}

	c.items[key] = c.queue.PushFront(ci)

	if c.queue.Len() > c.capacity {
		cacheItemKey, ok := c.queue.Back().Value.(*cacheItem)
		if !ok {
			return false
		}
		delete(c.items, cacheItemKey.key)
		c.queue.Remove(c.queue.Back())
	}

	return false
}

func (c *lruCache) Clear() {
	c.Lock()
	defer c.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) LruCache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
