// Package memo provides memoization of a function of type Func.
// Concurrency-safe, faster than memo2, but can duplicate calls to f(key).
package memo3

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
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)

		// Between the 2 critical sections, several goroutines may race to compute f(key)
		// and update the map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
