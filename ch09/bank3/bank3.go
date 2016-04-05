package bank3

import "sync"

var (
	mutex sync.Mutex
	balance int
)

func Deposit(amount int) {
	mutex.Lock()
	balance = balance + amount
	// BobK:  Why not defer?
	mutex.Unlock()
}

func Balance() int {
	mutex.Lock()
	b := balance
	// BobK:  Why not defer?
	mutex.Unlock()
	return b
}
