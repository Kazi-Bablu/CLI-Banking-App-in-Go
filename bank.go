package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFormFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("failed to read find balance file.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored balance value.")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance, err = getBalanceFormFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("-------")
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")

	for i := 0; i < 2; i++ {
		fmt.Println("What do you want to do ?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit memory")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
			fmt.Println("Your deposit:")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				//return
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else if choice == 3 {
			fmt.Println("Withdraw amount:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				// return
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid amount. You can't withdraw more than you have.")
				// return
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		} else {
			fmt.Println("Goodbye!")
			break
		}
	}
	fmt.Println("Thank for choosing our bank")

}
