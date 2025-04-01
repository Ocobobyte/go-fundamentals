package account

import (
	"fmt"
)

// Account struct represents a bank account.
type Account struct {
	Owner   string
	Balance float64
}

// Global default account initialized in init.
var DefaultAccount Account

// Deposit adds an amount to the account balance.
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
	fmt.Printf("Deposited: %.2f, New Balance: %.2f\n", amount, a.Balance)
}

// Withdraw deducts an amount from the account if sufficient funds exist.
func (a *Account) Withdraw(amount float64) bool {
	if amount > a.Balance {
		fmt.Println("Insufficient funds")
		return false
	}
	a.Balance -= amount
	fmt.Printf("Withdrawn: %.2f, New Balance: %.2f\n", amount, a.Balance)
	return true
}

// ApplyTransactions accepts a variable number of transactions.
// Positive values deposit money, while negative values trigger a withdrawal.
func (a *Account) ApplyTransactions(transactions ...float64) {
	for _, t := range transactions {
		if t >= 0 {
			a.Deposit(t)
		} else {
			a.Withdraw(-t)
		}
	}
}

// init runs before main and initializes the default account.
func init() {
	fmt.Println("Welcome to the Bank Account App!")
	// Creative initialization: setting up the default account with a welcome note.
	DefaultAccount = Account{Owner: "Alice", Balance: 1000.0}
	fmt.Printf("Initializing account for %s with a starting balance of %.2f\n", DefaultAccount.Owner, DefaultAccount.Balance)
}
