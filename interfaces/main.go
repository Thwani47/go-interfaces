package main

import "fmt"

type IBankAccount interface {
	GetBalance() int // 100 == 1 dollar
	Deposit(amount int)
	Withdraw(amount int) error
}

func main() {
	myAccounts := []IBankAccount{
		NewWellsFargo(),
		NewBitCoinAccount(),
	}

	for _, account := range myAccounts {
		account.Deposit(500)
		if err := account.Withdraw(70); err != nil {
			fmt.Printf("Error: %s\n", err)
		}
		balance := account.GetBalance()
		fmt.Printf("Balance: %d\n", balance)
	}

}
