package main

import (
	"fmt"
)

const accountBalanceFileName = "balance.txt"

func main() {

	var accountBalance, balanceError = ReadFloatFromFile(accountBalanceFileName)

	if balanceError != nil {
		presentError(balanceError)
	}

	fmt.Println("Welcome to your Banking App!")

	for {

		presentUserOptions()

		var choice int

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			accountBalance += deposit()
			fmt.Println("Your new balance: ", accountBalance)
			WriteFloatToFile(accountBalance, accountBalanceFileName)
			continue
		case 3:
			accountBalance -= withdraw(accountBalance)
			fmt.Println("Your new balance: ", accountBalance)
			WriteFloatToFile(accountBalance, accountBalanceFileName)
			continue
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for using our bank.")
			return
		}
	}
}

func deposit() float64 {
	var depositAmount float64
	fmt.Println("Deposit amount")
	fmt.Scan(&depositAmount)

	if depositAmount <= 0 {
		fmt.Println("Invalid amount. Must be greater than 0")
		return 0
	}
	return depositAmount
}

func withdraw(balance float64) float64 {
	var withdrawalAmount float64
	fmt.Println("Withdrawal amount")
	fmt.Scan(&withdrawalAmount)

	if withdrawalAmount > balance {
		fmt.Println("Not enough funds.")
		return 0
	}

	return withdrawalAmount
}
