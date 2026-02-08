package main

import (
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balanceFile.txt"

func main() {
	fmt.Println("Welcome to Go Bank!")
	for {
		choice := promptChoice()
		userChoice(choice)
		if choice == 4 {
			break
		}
	}
}

func promptChoice() (choice int) {
	fmt.Println("How may I help?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
	fmt.Print("Choice: ")
	fmt.Scan(&choice)

	return choice
}

func getBalance() string {
	data, _ := os.ReadFile(balanceFile)
	balance := string(data)
	fmtBalance := fmt.Sprintf("$%s is current balance\n\n", balance)
	return fmtBalance
}

func depositMoney(money float64) {
	if money < 0 {
		fmt.Print("Deposit must be $0 or greater\n\n")
		return
	}
	data, _ := os.ReadFile(balanceFile)
	balance := string(data)
	numericalBalance, _ := strconv.ParseFloat(balance, 64)
	balance = strconv.FormatFloat((numericalBalance + money), 'f', 2, 64)
	os.WriteFile(balanceFile, []byte(balance), 0644)
	fmt.Printf("$%.2f deposited\n$%s is new balance.\n\n", money, balance)
}

func withdrawMoney(money float64) {
	data, _ := os.ReadFile(balanceFile)
	balance := string(data)
	numericalBalance, _ := strconv.ParseFloat(balance, 64)
	if money > numericalBalance {
		fmt.Print("Withdraw exceeds available funds\n\n")
		return
	}
	balance = strconv.FormatFloat((numericalBalance - money), 'f', 2, 64)
	os.WriteFile(balanceFile, []byte(balance), 0644)
	fmt.Printf("$%.2f withdrawn\n$%s is new balance.\n\n", money, balance)
}

func userChoice(choice int) {
	var amount float64
	switch choice {
	case 1:
		fmt.Print(getBalance())
	case 2:
		fmt.Print("How much to deposit?: $")
		fmt.Scan(&amount)
		depositMoney(amount)
	case 3:
		fmt.Print("How much to withdraw?: $")
		fmt.Scan(&amount)
		withdrawMoney(amount)
	case 4:
		fmt.Println("Goodbye")
	}
}
