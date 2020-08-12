package main

import (
	"fmt"

	"github.com/YoonsooChang/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	dictionary.Add("hello", "world")
	def, err := dictionary.Search("hello")
	if err == nil {
		fmt.Println(def)
	} else {
		fmt.Println(err)
	}
	def, err = dictionary.Update("hello", "WALDOO")
	if err == nil {
		fmt.Println(def)
	} else {
		fmt.Println(err)
	}
	def, err = dictionary.Update("bye", "later")
	if err == nil {
		fmt.Println(def)
	} else {
		fmt.Println(err)
	}
	err = dictionary.Delete("Say")
	if err != nil {
		fmt.Println(err)
	}
}

/* ACCOUNTS TEST CODE
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
*/
