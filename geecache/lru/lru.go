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
		kv := ele.Value.(*entry) // Value means the value in the list element
		return kv.value, true
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() // find the lowest frequent one
	if ele != nil {    // if not null
		c.ll.Remove(ele) // remove it
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)                                      // delete from the hashmap
		c.currentBytes -= int64(len(kv.key)) + int64(kv.value.Len()) // minus the len
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value) // call back
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok { // exist key, since visit -> movetofront
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.currentBytes += int64(value.Len()) - int64(kv.value.Len()) // (new one) - (old one)
		kv.value = value
	} else { // do not exist
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.currentBytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.currentBytes {
		c.RemoveOldest() // if add the new node cause the bytes exceed the maximum value, we should remove the oldest until have enough space
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
