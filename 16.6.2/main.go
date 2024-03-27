package main

import (
	"fmt"
	"sync"
)

var _ BankClient = Account{} // это трюк для проверки типа: до тех пор пока Account не будет реализовывать интерфейс BankClient, программа не запустится

type BankClient interface {
	// Deposit deposits given amount to clients account
	Deposit(amount int)

	// Withdrawal withdraws given amount from clients account.
	// return error if clients balance less the withdrawal amount
	Withdrawal(amount int) error

	// Balance returns clients balance
	Balance() int
}

// Структура для имплементации заданного интерфейса
type Account struct {
	bMutex  *sync.RWMutex
	balance int
}

// Balance ...
func (a Account) Balance() int {
	return 0
}

// Deposit ...
func (a Account) Deposit(amount int) {

}

// Withdrawal ..
func (a Account) Withdrawal(amount int) error {
	return nil
}

// NewAccount ..
func NewAccount() Account {
	return Account{
		bMutex:  &sync.RWMutex{},
		balance: 0,
	}
}
func main() {
	account := NewAccount()
	fmt.Println(account.balance)
}
