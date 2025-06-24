//Encapsulations means restrcting access to internal state or behaviour

package main

import "fmt"

type Account struct {
	balance float64 // now here this balance is protected and cant be used from other package
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance = a.balance + amount // if someone wants to add balanace the only can use Deposite fun to add
	}
}

func (a *Account) BalannceCheck() float64 {
	return a.balance
}

func main() {
	a := &Account{balance: 2000}
	a.Deposit(10000)
	balance := a.BalannceCheck()

	fmt.Println(balance)
}
