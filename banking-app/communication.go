package main

import "fmt"

func presentUserOptions() {
	fmt.Println("What would you like to do?")
	fmt.Println("1: Check balance")
	fmt.Println("2: Deposit money")
	fmt.Println("3: Withdraw money")
	fmt.Println("4: Exit")
}

func presentError(error error) {
	fmt.Println("---------")
	fmt.Println("ERROR")
	fmt.Println(error)
	fmt.Println("---------")
}
