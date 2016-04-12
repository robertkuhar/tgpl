// Package memo provides memoization of a function of type Func.
// Concurrency-safe, but slow due to holding the lock for all accesses.
package memo2

import "sync"

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}


// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (value interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	// BobK:  why not defer?
	memo.mu.Unlock()
	return res.value, res.err
}
