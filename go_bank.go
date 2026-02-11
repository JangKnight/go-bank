package main

import (
	"errors"
	"fmt"
	"slices"

	"example.com/go-bank/accounts"
)

func main() {
	var account string
	fmt.Println("Welcome to Go Bank!")
	fmt.Print("Please enter your account filename: ")
	fmt.Scanln(&account)
	for {
		choice, _ := promptChoice()
		userChoice(choice, account)
		if choice == 5 {
			break
		}
	}
}

func promptChoice() (choice int, error error) {
	fmt.Println("How may I help?")
	fmt.Println("1. Create new account")
	fmt.Println("2. Check balance")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Withdraw Money")
	fmt.Println("5. Exit")
	fmt.Print("Choice: ")
	fmt.Scanln(&choice)

	if slices.Contains([]int{1, 2, 3, 4, 5}, choice) {
		return choice, nil
	}

	return 0, errors.New("Invalid choice")
}

func userChoice(choice int, account string) {
	var amount float64
	switch choice {
	case 0:
		fmt.Print("Invalid choice\n\n")
	case 1:
		accounts.CreateAccount(account)
	case 2:
		fmt.Print(accounts.GetBalance(account))
	case 3:
		fmt.Print("How much to deposit?: $")
		fmt.Scan(&amount)
		accounts.DepositMoney(amount, account)
	case 4:
		fmt.Print("How much to withdraw?: $")
		fmt.Scan(&amount)
		accounts.WithdrawMoney(amount, account)
	case 5:
		fmt.Println("Goodbye")
	}
}
