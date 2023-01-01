package account

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) { // convenstion object low first char
	a.balance += amount
}

func (a Account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("can't withdraw")

// Withdraw is minus
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
