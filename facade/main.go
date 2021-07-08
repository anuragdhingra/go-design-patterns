package main

import "fmt"

type facade struct {
	account *account
}
type Facade interface {
	AddMoney(accountID, amount int) bool
}
func (f facade) AddMoney(accountID, amount int) bool {
	// Any complex logic
	if f.account.IsValid(accountID) {
		return true
	}
	fmt.Println("amount to add is ", amount)
	return true
}

type account struct {}
func (a account) IsValid(accountID int) bool {
	return true
}

func main() {
	facadeInstance := facade{
		account: &account{},
	}

	if facadeInstance.AddMoney(123, 100) {
		fmt.Println("Money added")
	}
}
