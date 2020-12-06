package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin represents a numerical bitcoin value
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet represents a bitcoin wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit allows a bitcoin deposit into the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance gets the overall balance of the Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw allows a bitcoin withdrawal from the Wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}
	w.balance -= amount
	return nil
}
