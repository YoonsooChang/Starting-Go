package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

type reqRes struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan reqRes)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, c chan<- reqRes) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- reqRes{url: url, status: status}
}

/* MYDICT TEST CODE
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
*/

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
