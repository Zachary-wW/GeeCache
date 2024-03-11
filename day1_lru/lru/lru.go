package lru

import "container/list"

// import "container/list"

type Cache struct {
	maxBytes     int64 // maximum usage memory
	currentBytes int64 // current usage memory
	ll           *list.List
	cache        map[string]*list.Element      // the key is string, value is the element point of the doublelist
	OnEvicted    func(key string, value Value) // call back func
}

type entry struct {
	key   string
	value Value
} // the data structure in list value

type Value interface {
	Len() int // return the memory size
}

/*
the return value of this func is *Cache
*/
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// The Get method is (lru.Cache).Get
// func input is the string type: key
func (c *Cache) Get(key string) (value Value, ok bool) {
	ele, ok := c.cache[key] // find the element, but what about ok?
	if ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry) // what is happening here
		return kv.value, true
	}
	return
}
