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
		choice, err := promptChoice()
		userChoice(choice, account)
		if err != nil {
			fmt.Println(err)
		}
		if choice == 4 {
			break
		}
	}
}

func promptChoice() (choice int, error error) {
	fmt.Println("How may I help?")
	fmt.Println("0. Create new account")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Print("Choice: ")
	fmt.Scanln(&choice)

	if slices.Contains([]int{0, 1, 2, 3, 4}, choice) {
		return choice, nil
	}

	return 0, errors.New("Invalid choice")
}

func userChoice(choice int, account string) {
	var amount float64
	switch choice {
	case 0:
		accounts.CreateAccount(account)
	case 1:
		fmt.Print(accounts.GetBalance(account))
	case 2:
		fmt.Print("How much to deposit?: $")
		fmt.Scan(&amount)
		accounts.DepositMoney(amount, account)
	case 3:
		fmt.Print("How much to withdraw?: $")
		fmt.Scan(&amount)
		accounts.WithdrawMoney(amount, account)
	case 4:
		fmt.Println("Goodbye")
	}
}
