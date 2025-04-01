package main

import (
	"fmt"

	"github.com/ocobobyte/bank-account/account"
)

func main() {
	// Defer ensures this message prints last, after main completes.
	defer fmt.Println("Transaction complete. Thank you!")

	fmt.Printf("Account Owner: %s, Starting Balance: %.2f\n", account.DefaultAccount.Owner, account.DefaultAccount.Balance)

	// Perform individual transactions.
	account.DefaultAccount.Deposit(200)
	account.DefaultAccount.Withdraw(150)

	// Process multiple transactions at once using a variadic function.
	account.DefaultAccount.ApplyTransactions(100, -50, 300, -600)

	fmt.Printf("Final Balance for %s: %.2f\n", account.DefaultAccount.Owner, account.DefaultAccount.Balance)
}
