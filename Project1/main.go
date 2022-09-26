package main

import (
	"fmt"

	"github.com/benloonshots/learnGo/Project1/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	fmt.Println(account)

}
