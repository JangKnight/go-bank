package accounts

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func getAccounts() []string {
	files, _ := os.ReadDir(".")
	var accounts []string
	for _, file := range files {
		if !file.IsDir() {
			accounts = append(accounts, file.Name())
		}
	}
	return accounts
}

func CreateAccount(accountName string) {
	accounts := getAccounts()
	if slices.Contains(accounts, accountName) {
		fmt.Printf("Account %s already exists\n\n", accountName)
	} else {
		os.WriteFile(accountName, []byte("0.00"), 0644)
		fmt.Printf("Account %s created\n\n", accountName)
	}
}

func GetBalance(balanceFile string) string {
	data, _ := os.ReadFile(balanceFile)
	balance := string(data)
	fmtBalance := fmt.Sprintf("$%s is current balance\n\n", balance)
	return fmtBalance
}

func DepositMoney(money float64, balanceFile string) {
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

func WithdrawMoney(money float64, balanceFile string) {
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
