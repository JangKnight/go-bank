package accounts

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func getAccounts() []string {
	files, _ := os.ReadDir("./accounts")
	var accounts []string
	for _, file := range files {
		if !file.IsDir() {
			accounts = append(accounts, file.Name())
		}
	}
	return accounts
}

func accountExists(accountName string) (exists bool) {
	accounts := getAccounts()
	exists = false
	if slices.Contains(accounts, accountName) {
		exists = true
	}
	return exists
}

func CreateAccount(accountName string) {
	if accountExists(accountName) {
		fmt.Printf("Account %s already exists\n\n", accountName)
	} else {
		os.WriteFile("./accounts/"+accountName, []byte("0.00"), 0644)
		fmt.Printf("Account %s created\n\n", accountName)
	}
}

func GetBalance(accountName string) (fmtBalance string) {
	if accountExists(accountName) {
		data, _ := os.ReadFile("./accounts/" + accountName)
		balance := string(data)
		fmtBalance = fmt.Sprintf("$%s is current balance\n\n", balance)
	} else {
		fmt.Printf("Account %s does NOT exist!\n\n", accountName)
	}
	return fmtBalance
}

func DepositMoney(money float64, accountName string) {
	if accountExists(accountName) {
		if money < 0 {
			fmt.Print("Deposit must be $0 or greater\n\n")
			return
		}
		data, _ := os.ReadFile("./accounts/" + accountName)
		balance := string(data)
		numericalBalance, _ := strconv.ParseFloat(balance, 64)
		balance = strconv.FormatFloat((numericalBalance + money), 'f', 2, 64)
		os.WriteFile("./accounts/"+accountName, []byte(balance), 0644)
		fmt.Printf("$%.2f deposited\n$%s is new balance.\n\n", money, balance)
	} else {
		fmt.Printf("Account %s does NOT exist!\n\n", accountName)
	}
}

func WithdrawMoney(money float64, accountName string) {
	if accountExists(accountName) {
		data, _ := os.ReadFile("./accounts/" + accountName)
		balance := string(data)
		numericalBalance, _ := strconv.ParseFloat(balance, 64)
		if money > numericalBalance {
			fmt.Print("Withdraw exceeds available funds\n\n")
			return
		}
		balance = strconv.FormatFloat((numericalBalance - money), 'f', 2, 64)
		os.WriteFile("./accounts/"+accountName, []byte(balance), 0644)
		fmt.Printf("$%.2f withdrawn\n$%s is new balance.\n\n", money, balance)
	} else {
		fmt.Printf("Account %s does NOT exist!\n\n", accountName)
	}
}
