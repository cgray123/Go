package sub

import (
	"fmt"
	"sort"
	"time"
)

type BankAccount struct {
	AccountBalance float64
	AccountType    string
	AccountID      int
	AccountHolder  []string
	ActionDate     []string
	bank           *Bank
}
type Bank struct {
	Accounts []BankAccount
	BankAccount
}

func (b *BankAccount) SetAccBal(x float64) {
	b.AccountBalance = x
}
func (b *BankAccount) GetAccBal() {
	fmt.Println(b.AccountBalance)
}
func (b *BankAccount) GetActDate() {
	fmt.Println(b.ActionDate)
}
func (b *BankAccount) SetAccType(x string) {
	b.AccountType = x
}
func (b *BankAccount) SetAccHolder(x []string) {
	b.AccountHolder = x
}
func (b *BankAccount) SetAccID(x int) {
	b.AccountID = x
}
func (b *BankAccount) _setActDate(x string) {
	b.ActionDate = append(b.ActionDate, x)
}
func (b *Bank) _addToBank(x BankAccount) {
	b.Accounts = append(b.Accounts, x)
}
func (b *Bank) _updateBank(z BankAccount) {
	b.Accounts = append(b.Accounts[:z.AccountID], b.Accounts[z.AccountID+1:]...)
	b._addToBank(z)
}

func (b *BankAccount) Withdraw(x float64) {
	if b.AccountBalance < x {
		fmt.Println("Can not withdraw more than what is in the account")
	} else if x <= 0 {
		fmt.Println("Please enter an amount greater than zero")
	} else {
		b.SetAccBal(b.AccountBalance - x)
		data := " ID:" + fmt.Sprint(b.AccountID) + "  withdraw " + time.Now().String() + " Withdraw amount: " + fmt.Sprint(x) + " |"
		b._setActDate(data)
		b.bank._updateBank(*b)
	}

}
func (b *BankAccount) Deposit() {

	fmt.Println("please input deposit amount")
	var x float64
	fmt.Scanln(&x)
	if x <= 0 {
		fmt.Println("Please input a number greater than 0")
	} else {
		b.SetAccBal(b.AccountBalance + x)
		data := "ID:" + fmt.Sprint(b.AccountID) + " deposit " + time.Now().String() + " Deposit amount: " + fmt.Sprint(x) + " |"
		b._setActDate(data)
		b.bank._updateBank(*b)
	}
}
func (b *BankAccount) MakeAccount(z *Bank) {
	b.bank = z
	fmt.Println("input starting balance")
	var x float64
	fmt.Scanln(&x)
	b.SetAccBal(x)
	fmt.Println("input first and last name")
	var y, e string
	fmt.Scanln(&y, &e)
	b.SetAccHolder([]string{y, e})
	fmt.Println("input account type")
	var o string
	fmt.Scanln(&o)
	b.SetAccType(o)
	if len(z.Accounts) == 0 {
		b.SetAccID(0)
	} else {
		b.SetAccID(b.bank.Accounts[cap(b.bank.Accounts)-1].AccountID + 1)
	}

	data := "ID:" + fmt.Sprint(b.AccountID) + " creation " + time.Now().String() + " Starting Balance: " + fmt.Sprint(x) + " |"
	b._setActDate(data)

	z._addToBank(*b)

}
func (b *Bank) SortName() {
	var temp []string
	for i := 0; i < len(b.Accounts); i++ {
		temp = append(temp, b.Accounts[i].AccountHolder[0])
	}
	sort.Strings(temp)
	fmt.Println(temp)
}
func (b *Bank) SearchByID(x int) {
	fmt.Println("Account type:", b.Accounts[x].AccountType, "Balance:", b.Accounts[x].AccountBalance, "Name:", b.Accounts[x].AccountHolder, "All account actions:", b.Accounts[x].ActionDate)
}
