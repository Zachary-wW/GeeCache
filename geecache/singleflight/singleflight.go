package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val any
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// no matter how many times the Do call, the fn just excute once for the same key
func (g *Group) Do(key string, fn func() (any, error)) (any, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()         // if the request is running, just wait
		return c.val, c.err // request over return results
	}
	c := new(call)
	c.wg.Add(1) // If the counter becomes zero, all goroutines blocked on Wait are released.
	g.m[key] = c
	g.mu.Unlock()

	c.val, c.err = fn() // call fn
	c.wg.Done()         // over request -1

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
