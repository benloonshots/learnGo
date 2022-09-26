package main

import (
	"fmt"

	"github.com/benloonshots/learnGo/Project1/banking"
)

func main() {
	account := banking.Account{Owner: "nico", Balance: 10000}
	fmt.Println(account)
}
