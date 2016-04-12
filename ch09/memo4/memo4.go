// Package memo provides memoization of a function of type Func.
// Concurrency-safe,
package memo4

import "sync"

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // close when res is ready
}


// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

type Memo struct {
	f     Func
	my    sync.Mutex
	cache map[string]*entry
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		// This is a repeat request for this key.
		memo.mu.Unlock()
		<-e.ready
	}
	memo.mu.Unlock()
	return e.res.value, e.res.err
}
