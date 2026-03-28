package account

import (
	"sync"
)

// Define the Account type here.
type Account struct {
	balance int64
	open    bool
	mu      *sync.Mutex
}

// Open creates a new account with the given initial balance.
// If the initial balance is negative, it returns nil.
func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, open: true, mu: &sync.Mutex{}}
}

func (account *Account) Balance() (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open {
		return 0, false
	}
	return account.balance, true
}

func (account *Account) Deposit(amount int64) (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open || account.balance+amount < 0 {
		return 0, false
	}
	account.balance += amount
	return account.balance, true
}

func (account *Account) Close() (int64, bool) {
	account.mu.Lock()
	defer account.mu.Unlock()
	if !account.open {
		return 0, false
	}
	balance := account.balance
	account.balance = 0
	account.open = false
	return balance, true
}
