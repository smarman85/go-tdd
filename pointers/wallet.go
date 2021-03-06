package main

import (
  "fmt"
  "errors"
)

type Bitcoin int

type Stringer interface {
  String() string
}

type Wallet struct {
  balance Bitcoin
}

func (w *Wallet) Deposit(ammount Bitcoin) {
  //fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
  w.balance += ammount
}

func (w *Wallet) Balance() Bitcoin {
  return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
  if amount > w.balance {
    return ErrInsufficientFunds
  }
  w.balance -= amount
  return nil
}

func (b Bitcoin) String() string {
  return fmt.Sprintf("%d BTC", b)
}
