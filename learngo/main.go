package main

import (
	"fmt"

	"github.com/YoonsooChang/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Yoonsoo")
	account.Deposit(1000)
	err := account.Withdraw(2000)
	if err != nil {
		// log.Fatalln(err)
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)

}
