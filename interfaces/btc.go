package main

import "errors"

type BitCoinAccount struct {
	balance int
	fee     int
}

func NewBitCoinAccount() *BitCoinAccount {
	return &BitCoinAccount{
		balance: 0,
		fee:     300,
	}
}

func (bc *BitCoinAccount) GetBalance() int {
	return bc.balance
}

func (bc *BitCoinAccount) Deposit(amount int) {
	bc.balance += amount
}

func (bc *BitCoinAccount) Withdraw(amount int) error {
	newBalance := bc.balance - bc.fee - amount
	if newBalance < 0 {
		return errors.New("insufficient funds")
	}

	bc.balance = newBalance
	return nil
}
