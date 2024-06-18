package main

import "errors"

type WellsFargo struct {
	balance int
}

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{
		balance: 0,
	}
}

func (wf *WellsFargo) GetBalance() int {
	return wf.balance
}

func (wf *WellsFargo) Deposit(amount int) {
	wf.balance += amount
}

func (wf *WellsFargo) Withdraw(amount int) error {
	if wf.balance < amount {
		return errors.New("insufficient funds")
	}

	wf.balance -= amount
	return nil
}
