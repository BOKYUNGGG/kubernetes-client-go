package main

import (
	"fmt"

	"github.com/SEOBOKYUNGGG/kubernetes-client-go/banking"
)

func main() {
	account := banking.Account{Owner: "nico", Balance: 1000}
	fmt.Println(account)
}
