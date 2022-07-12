package main

import (
	"fmt"
	"os"
	"strconv"
)
// banking mod 6 inputs are given as you run the program
func main() {
	B := Bank{Wallets: make([]*Wallet, 0)}
	s := AccOwner{ID: 0, Address: "123", Type: "individual"}
	b := BankAcc{AccNum: "156054", Balance: 100, AccType: "checking", Owner: s}
	b2 := BankAcc{AccNum: "15604", Balance: 1001, AccType: "investment", Owner: s}
	w := Wallet{WalletID: "aoeu", Accounts: []*BankAcc{&b, &b2}, WalletOwner: s, HostBank: &B}
	s1 := AccOwner{ID: 1, Address: "123", Type: "individual"}
	b1 := BankAcc{AccNum: "156", Balance: 100, AccType: "checking", Owner: s1}
	w2 := Wallet{WalletID: "abc", Accounts: []*BankAcc{&b1}, WalletOwner: s1, HostBank: &B}
	B.Wallets = append(B.Wallets, &w)
	B.Wallets = append(B.Wallets, &w2)

	b.CalInterest()
	b2.CalInterest()
	b1.CalInterest()
	w.Terminal()
}

type BankAcc struct {
	AccNum       string
	Balance      float64
	AccType      string
	Owner        AccOwner
	InterestRate float64
}
type AccOwner struct {
	ID      int
	Address string
	Type    string
}
type Wallet struct {
	WalletID    string
	Accounts    []*BankAcc
	WalletOwner AccOwner
	HostBank    *Bank
}
type Bank struct {
	Wallets []*Wallet
}

//takes in a float and adds it to the balance of the given bank account
func (b *BankAcc) Deposit(x float64) {
	fmt.Scanln(&x)
	if x <= 0 {
		fmt.Println("Please input a number greater than zero")
	} else {
		b.Balance += x
	}
}

//takes in a float and subtract it to the balance of the given bank account
func (b *BankAcc) Withdraw(x float64) {
	if b.Balance < x {
		fmt.Println("Can not withdraw more than what is in the account")
	} else if x <= 0 {
		fmt.Println("Please enter an amount greater than zero")
	} else {
		b.Balance -= x
	}
}

//applies the interest formula and updates the balance
func (b *BankAcc) ApplyInterest() {
	b.Balance = b.Balance * (1 + b.InterestRate)
}

//the account that this method is called on is the scr, the account that is passed as a perm is the target
func (b *BankAcc) Wire(b2 *BankAcc, x float64) {
	if b.Balance < x {
		fmt.Println("Balance too small for this wire")
	} else {
		b.Balance -= x
		b2.Balance += x
	}
}

//returns the bankacc obj
func (b *BankAcc) Display() BankAcc {
	return *b
}

//changes the address of the owner
func (a *AccOwner) ChangeAddress(x string) {
	a.Address = x
}

//groups all accounts by type and displays them
func (w *Wallet) DisplayAcc() {
	var c, in, s []BankAcc
	for i := 0; i < len(w.Accounts); i++ {
		if w.Accounts[i].AccType == "checking" {
			c = append(c, w.Accounts[i].Display())
		} else if w.Accounts[i].AccType == "investment" {
			in = append(in, w.Accounts[i].Display())
		} else if w.Accounts[i].AccType == "saving" {
			s = append(s, w.Accounts[i].Display())
		}
	}
	fmt.Println("Checking Accounts:", c)
	fmt.Println("Investment Accounts:", in)
	fmt.Println("Saving Accounts:", s)
}

//prints all balances in the wallet and prints the sum
func (w *Wallet) DisplayBals() {
	var tot int
	for i := 0; i < len(w.Accounts); i++ {
		fmt.Println(w.Accounts[i].AccNum, "-", w.Accounts[i].Balance)
		tot += int(w.Accounts[i].Balance)
	}
	fmt.Println("Total Balance:", tot)
}

//takes in source then target then amount and subtracts the amount from the source and adds the amount to the target
func (w *Wallet) Wire(b, b2 *BankAcc, x float64) {
	if b.Balance < x {
		fmt.Println("Balance too small for this wire")
		flag := true
		for i := 0; i < len(w.Accounts); i++ {
			if w.Accounts[i].Balance >= x {
				fmt.Println("You have the funds in account", w.Accounts[i].AccNum)
				flag = false
			}

		}
		if flag {
			fmt.Println("None of your Accounts have the funds to do this wire")
		}
	} else {
		b.Balance -= x
		b2.Balance += x
	}
}

//applies interest to all accounts in the wallet
func (w *Wallet) ApplyInterest() {
	for i := 0; i < len(w.Accounts); i++ {
		w.Accounts[i].Balance = w.Accounts[i].Balance * (1 + w.Accounts[i].InterestRate)
	}

}

//sets interest rate based on type
func (b *BankAcc) CalInterest() {
	var in float64
	if b.AccType == "checking" {
		in = .005
	} else if b.AccType == "investment" {
		in = .01
	} else if b.AccType == "saving" {
		in = .02
	}
	if b.Owner.Type == "individual" {
		if b.AccType == "saving" {
			b.InterestRate = in + .03
		} else {
			b.InterestRate = in + in
		}
	} else {
		b.InterestRate = in
	}

}

//main menu passes you off to other menu methods
func (w *Wallet) Terminal() {
	fmt.Println("Input 1 to view all Accounts")
	fmt.Println("Input 2 to interact with an Account")
	fmt.Println("Input 3 to interact with your wallet")
	fmt.Println("Input y to exit")
	var input string
	fmt.Scanln(&input)
	x, _ := strconv.Atoi(input)
	if x == 1 {
		w.DisplayAcc()
		w.Terminal()
	} else if x == 2 {
		fmt.Println("Input the ID of the account you would like to interact with, all IDs are show below")
		w.PrintIDs()
		var num string
		fmt.Scanln(&num)
		for i := 0; i < len(w.Accounts); i++ {
			if w.Accounts[i].AccNum == num {
				w.Accounts[i].InAccount(w)

			}
		}
		fmt.Println("You did not enter a vaild ID")
		w.Terminal()
	} else if x == 3 {
		w.InWallet()

	} else if input == "y" {
		os.Exit(3)
	} else {
		fmt.Println("please input one of the given options")
		w.Terminal()
	}
}

//account menu simple method calls
func (b *BankAcc) InAccount(i *Wallet) {
	fmt.Println("Input 1 to view Account")
	fmt.Println("Input 2 to withdraw from your Account")
	fmt.Println("Input 3 to deposit  from your Account")
	fmt.Println("Input 4 to apply interest to you Account")
	fmt.Println("Input y to return to the main menu")
	var w float64
	var input string
	fmt.Scanln(&input)
	x, _ := strconv.Atoi(input)
	if x == 1 {
		fmt.Println(b)
		b.InAccount(i)
	} else if x == 2 {
		fmt.Println("How much would you like to withdraw?")
		fmt.Scanln(&w)
		b.Withdraw(w)
		fmt.Println("Your new balance", b.Balance)
		b.InAccount(i)
	} else if x == 3 {
		fmt.Println("How much would you like to deposit?")
		fmt.Scanln(&w)
		b.Deposit(w)
		fmt.Println("Your new balance", b.Balance)
		b.InAccount(i)
	} else if x == 4 {
		b.ApplyInterest()
		fmt.Println("Your new balance", b.Balance)
		b.InAccount(i)
	} else if input == "y" {
		i.Terminal()
	} else {
		fmt.Println("You did not enter a vaild option")
		b.InAccount(i)
	}

}

//wallet menu helper logic for wire
func (w *Wallet) InWallet() {
	fmt.Println("Input 1 display all account balances")
	fmt.Println("Input 2 make a wire")
	fmt.Println("Input 3 to apply interest to all Accounts in you wallet")
	fmt.Println("Input y to return to the main menu")
	var input string
	fmt.Scanln(&input)
	x, _ := strconv.Atoi(input)
	if x == 1 {
		w.DisplayBals()
		w.InWallet()
	} else if x == 2 {
		fmt.Println("Input the ID of the account your are sending from,all IDs are show below")
		w.PrintIDs()
		var num string
		var sB, tB BankAcc
		fmt.Scanln(&num)
		for i := 0; i < len(w.Accounts); i++ {
			if w.Accounts[i].AccNum == num {
				sB = *w.Accounts[i]
			}
		}
		if sB.AccNum == "" {
			fmt.Println("You did not enter a vaild ID")
			w.InWallet()
		} else {
		retry:
			fmt.Println("Is this an internal or external wire (input i or e)")
			var in string
			fmt.Scanln(&in)
			if in == "i" {
				fmt.Println("Input the ID of the account your are sending to,all IDs are show below")
				w.PrintIDs()
				fmt.Scanln(&num)
				for i := 0; i < len(w.Accounts); i++ {
					if w.Accounts[i].AccNum == num {
						tB = *w.Accounts[i]
					}
				}
				if tB.AccNum == "" {
					fmt.Println("You did not enter a vaild ID")
					goto retry
				} else {
					fmt.Println("input wire amount")
					var send float64
					fmt.Scanln(&send)
					w.Wire(&sB, &tB, send)
					fmt.Println("your new balances")
					fmt.Println(sB.Balance, tB.Balance)
					w.InWallet()
				}
			} else if in == "e" {
				fmt.Println("Input the ID of the account your are sending to,all IDs are show below")
				for i := 0; i < len(w.HostBank.Wallets); i++ {
					for j := 0; j < len(w.HostBank.Wallets[i].Accounts); j++ {
						if w.HostBank.Wallets[i].WalletID == w.WalletID {
							break
						}
						fmt.Println(w.HostBank.Wallets[i].Accounts[j].AccNum)
					}
				}
				fmt.Scanln(&num)
				for i := 0; i < len(w.HostBank.Wallets); i++ {
					for j := 0; j < len(w.HostBank.Wallets[i].Accounts); j++ {
						if w.HostBank.Wallets[i].WalletID == w.WalletID {
							break
						}
						if w.HostBank.Wallets[i].Accounts[j].AccNum == num {
							tB = *w.Accounts[i]
						}
					}
				}
				if tB.AccNum == "" {
					fmt.Println("You did not enter a vaild ID")
					goto retry
				} else {
					fmt.Println("input wire amount")
					var send float64
					fmt.Scanln(&send)
					w.Wire(&sB, &tB, send)
					fmt.Println("your new balances")
					fmt.Println(sB.Balance, tB.Balance)
					w.InWallet()
				}

			} else {
				fmt.Println("You did not enter a vaild option")
				goto retry
			}
		}

	} else if x == 3 {
		w.ApplyInterest()
		fmt.Println("These are your new balances")
		w.DisplayBals()
		w.InWallet()
	} else if input == "y" {
		w.Terminal()
	} else {
		fmt.Println("You did not enter a vaild option")
		w.InWallet()
	}

}

//prints all ids of a wallet
func (w *Wallet) PrintIDs() {
	for i := 0; i < len(w.Accounts); i++ {
		fmt.Println(w.Accounts[i].AccNum)
	}
}
