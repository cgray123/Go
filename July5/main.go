package main

import (
	"mymod/sub"
)

func main() {

	b := new(sub.Bank)
	a := new(sub.BankAccount)
	a.MakeAccount(b)
	a.GetAccBal()
	a.Deposit()
	a.GetAccBal()
	a.Withdraw(90)
	a.GetAccBal()
	a.GetActDate()

	c := new(sub.BankAccount)
	c.MakeAccount(b)
	c.GetActDate()

	d := new(sub.BankAccount)
	d.MakeAccount(b)
	b.SortName()
	b.SearchByID(0)

}
