package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't Withdraw")

// NewAccount Creates account
func NewAccount(owner string) *account {
	_account := account{owner: owner, balance: 0}
	return &_account
}

// Deposit x Amount To Account
func (acc *account) Deposit(amount int) {
	acc.balance += amount
}

// Return balance Of Account
func (acc *account) Balance() int {
	return acc.balance
}

// Withdraw From Account
func (acc *account) Withdraw(amount int) error {
	if acc.balance < amount {
		return errNoMoney
	}
	acc.balance -= amount
	return nil
}

// Change Owner Of Account
func (acc *account) ChangeOwner(newOwner string) {
	acc.owner = newOwner
}

// Return Owner Of Account
func (acc *account) Owner() string {
	return acc.owner
}

func (acc *account) String() string {
	return fmt.Sprint(acc.Owner(), "'s account. \nHas : ", acc.Balance())
}
