package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	id      int
	balance int
	mu      sync.Mutex
}

func transfer(from, to *Account, amount int) {
	// Lock accounts in a consistent order to avoid deadlock.
	if from.id < to.id {
		from.mu.Lock()
		to.mu.Lock()
	} else {
		to.mu.Lock()
		from.mu.Lock()
	}
	defer from.mu.Unlock()
	defer to.mu.Unlock()

	fmt.Printf("[%d] locked from -> [%d]\n", from.id, to.id)
	time.Sleep(10 * time.Millisecond) // Simulate processing time

	from.balance -= amount
	to.balance += amount
}

func main() {
	bills := &Account{id: 1, balance: 1000}
	bilbo := &Account{id: 2, balance: 500}

	go transfer(bills, bilbo, 100)
	go transfer(bilbo, bills, 50)

	time.Sleep(2 * time.Second) // Wait for goroutines to finish

	fmt.Printf("Bills balance: %d\n", bills.balance)
	fmt.Printf("Bilbo balance: %d\n", bilbo.balance)

}
